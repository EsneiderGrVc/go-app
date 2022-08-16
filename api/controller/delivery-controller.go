package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/EsneiderGrVc/go_server/entity"
	"github.com/EsneiderGrVc/go_server/errors"
	"github.com/EsneiderGrVc/go_server/services"
	"github.com/EsneiderGrVc/go_server/validators"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type controller struct{}

var deliveryService services.DeliveryService

type DeliveryController interface {
	GetDelivery(w http.ResponseWriter, r *http.Request)
	GetAllDeliveries(w http.ResponseWriter, r *http.Request)
	CreateDelivery(w http.ResponseWriter, r *http.Request)
}

func NewDeliveryController(service services.DeliveryService) DeliveryController {
	deliveryService = service
	return &controller{}
}

// @Description Get a specific document by id.
// @Tags deliveries
// @Produce json
// @Param id path string true "delivery id"
// @Router /deliveries/{id} [get]
// @Success 200 {object} entity.Delivery
// @Failure 400
func (*controller) GetDelivery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	result, err := deliveryService.GetDeliveryById(vars["id"])
	if err != nil {
		log.Fatalf("Failed to get a document: %v", err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// @Description  Get all the documents inside of deliveries collection.
// @Tags         deliveries
// @Produce      json
// @Router       /deliveries [get]
// @Success 200 {array} entity.Delivery
// @Failure 400
func (*controller) GetAllDeliveries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	posts, err := deliveryService.GetAllDeliveries()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error getting the post"})
		return
	}

	w.WriteHeader(http.StatusOK)
	if len(posts) == 0 {
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "The collection has no documents"})
		return
	}
	json.NewEncoder(w).Encode(posts)
}

// @Description Create a new delivery
// @Tags deliveries
// @Produce json
// @Accept json
// @Param delivery body entity.PostDelivery true "delivery"
// @Router /deliveries [post]
// @Success 200 {object} entity.Delivery
// @Failure 400
func (*controller) CreateDelivery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var delivery entity.Delivery

	err := json.NewDecoder(r.Body).Decode(&delivery)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	if len(delivery.Id) == 0 {
		delivery.Id = uuid.Must(uuid.NewRandom()).String()
	}
	if len(delivery.CreationDate) == 0 {
		delivery.CreationDate = time.Now().Format("2006-01-02T15:04:05Z07:00")
	}

	vError := validators.NewValidator().ValidateDelivery(&delivery)
	if vError != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: vError.Error()})
		return
	}

	err1 := deliveryService.Validate(&delivery)
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := deliveryService.CreateDelivery(&delivery)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error saving the post"})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

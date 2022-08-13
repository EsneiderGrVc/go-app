package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/EsneiderGrVc/go_server/entity"
	"github.com/EsneiderGrVc/go_server/errors"
	"github.com/EsneiderGrVc/go_server/services"
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

func (*controller) CreateDelivery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var delivery entity.Delivery
	err := json.NewDecoder(r.Body).Decode(&delivery)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
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

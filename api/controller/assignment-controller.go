package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/EsneiderGrVc/go_server/entity"
	"github.com/EsneiderGrVc/go_server/errors"
	"github.com/EsneiderGrVc/go_server/services"
)

type assignmentController struct{}

var asgnService services.AssignmentService

type AssignmentController interface {
	AssignBotToDelivery(w http.ResponseWriter, r *http.Request)
}

func NewAssignmentController(service services.AssignmentService) AssignmentController {
	asgnService = service
	return &assignmentController{}
}

// @Description Assign a bot to an order.
// @Tags assignments
// @Produce json
// @Accept json
// @Param assignment body entity.RequestAssignment true "Assignment"
// @Router /assignments [post]
// @Success 200 {object} entity.Assignment
// @Failure 400
func (*assignmentController) AssignBotToDelivery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var asgn entity.Assignment

	err := json.NewDecoder(r.Body).Decode(&asgn)
	asgn.AssignmentDate = time.Now().Format("2006-01-02T15:04:05Z07:00")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	result, err := asgnService.AssignBotToDelivery(&asgn)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error assigning a bot"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

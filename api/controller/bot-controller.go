package controller

import (
	"encoding/json"
	"net/http"

	"github.com/EsneiderGrVc/go_server/entity"
	"github.com/EsneiderGrVc/go_server/errors"
	"github.com/EsneiderGrVc/go_server/services"
)

type botController struct{}

var botService services.BotService

type BotController interface {
	CreateBot(w http.ResponseWriter, r *http.Request)
	GetBotsOrderByZone(w http.ResponseWriter, r *http.Request)
}

func NewBotController(service services.BotService) BotController {
	botService = service
	return &botController{}
}

// @Description Create a new bot.
// @Tags bots
// @Produce json
// @Accept json
// @Param bot body entity.PostBot true "bot"
// @Router /bots [post]
// @Success 200 {object} entity.Bot
// @Failure 400
func (*botController) CreateBot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bot entity.Bot
	err := json.NewDecoder(r.Body).Decode(&bot)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	result, err2 := botService.CreateBot(&bot)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error saving the post"})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// @Description  Get all bots ordered by Zone.
// @Tags         bots
// @Produce      json
// @Router       /bots [get]
// @Success 200 {array} entity.Bot
// @Failure 400
func (*botController) GetBotsOrderByZone(w http.ResponseWriter, r *http.Request) {
	bots, err := botService.GetBotsOrderZone()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error getting bots"})
		return
	}

	w.WriteHeader(http.StatusOK)
	if len(bots) == 0 {
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "The collection has no documents"})
		return
	}
	json.NewEncoder(w).Encode(bots)
}

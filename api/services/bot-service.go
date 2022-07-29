package services

import (
	"github.com/EsneiderGrVc/go_server/entity"
	"github.com/EsneiderGrVc/go_server/repository"
)

type botServ struct{}

type BotService interface {
	CreateBot(bot *entity.Bot) (*entity.Bot, error)
	GetBotsOrderZone() ([]map[string]interface{}, error)
}

var botRepo repository.BotRepository

func NewBotService(r repository.BotRepository) BotService {
	botRepo = r
	return &botServ{}
}

func (*botServ) CreateBot(bot *entity.Bot) (*entity.Bot, error) {
	return botRepo.Save(bot)
}

func (*botServ) GetBotsOrderZone() ([]map[string]interface{}, error) {
	return botRepo.GetAll()
}

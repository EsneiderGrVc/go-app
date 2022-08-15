package validators

import (
	"github.com/EsneiderGrVc/go_server/entity"
)

type validate struct{}

type Validate interface {
	ValidateBot(bot *entity.Bot) error
}

func NewValidator() Validate {
	return &validate{}
}

func (*validate) ValidateBot(bot *entity.Bot) error {
	return nil
}

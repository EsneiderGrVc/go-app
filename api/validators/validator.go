package validators

import (
	"errors"
	"reflect"

	"github.com/EsneiderGrVc/go_server/entity"
)

type validate struct{}

type Validate interface {
	ValidateBot(bot *entity.Bot) error
	ValidateDelivery(delivery *entity.Delivery) error
}

func NewValidator() Validate {
	return &validate{}
}

func (*validate) ValidateBot(bot *entity.Bot) error {
	return nil
}

func (*validate) ValidateDelivery(delivery *entity.Delivery) error {
	v := reflect.ValueOf(delivery).Elem()
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == "" {
			return errors.New(v.Type().Field(i).Name + " attribute is empty")
		}
	}
	return nil
}

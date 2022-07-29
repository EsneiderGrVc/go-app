package services

import (
	"github.com/EsneiderGrVc/go_server/entity"
	"github.com/EsneiderGrVc/go_server/repository"
)

type service struct{}

type DeliveryService interface {
	Validate(delivery *entity.Delivery) error
	CreateDelivery(delivery *entity.Delivery) (*entity.Delivery, error)
	GetAllDeliveries() ([]map[string]interface{}, error)
	GetDeliveryById(id string) (map[string]interface{}, error)
	// DeleteDocument(post *entity.Delivery) (*entity.Delivery, error)
}

var deliveryRepo repository.DeliveryRepository

func NewDeliveryService(r repository.DeliveryRepository) DeliveryService {
	deliveryRepo = r
	return &service{}
}

func (*service) Validate(post *entity.Delivery) error {
	// if post == nil {
	// 	err := errors.New("The post is empty")
	// 	return err
	// }
	// if post.Title == "" {
	// 	err := errors.New("The post title is empty")
	// 	return err
	// }
	return nil
}

func (*service) GetDeliveryById(id string) (map[string]interface{}, error) {
	return deliveryRepo.GetDocumentById(id)
}

func (*service) CreateDelivery(delivery *entity.Delivery) (*entity.Delivery, error) {
	return deliveryRepo.Save(delivery)
}

func (*service) GetAllDeliveries() ([]map[string]interface{}, error) {
	return deliveryRepo.GetAll()
}

// func (*service) DeleteDocument(post *entity.Delivery) (*entity.Delivery, error) {
// 	return repo.DeleteDocument(post)
// }

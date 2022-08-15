package services

import (
	"github.com/EsneiderGrVc/go_server/entity"
	"github.com/EsneiderGrVc/go_server/repository"
)

type assignmentServ struct{}

type AssignmentService interface {
	AssignBotToDelivery(assignment *entity.Assignment) (*entity.Assignment, error)
}

var asgnRepo repository.AssigmentRepository

func NewAssignmentService(r repository.AssigmentRepository) AssignmentService {
	asgnRepo = r
	return &assignmentServ{}
}

func (*assignmentServ) AssignBotToDelivery(assignment *entity.Assignment) (*entity.Assignment, error) {
	return asgnRepo.AssignBotToDelivery(assignment)
}

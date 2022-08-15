package repository

import (
	"context"
	"errors"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/EsneiderGrVc/go_server/entity"
	"google.golang.org/api/option"
)

type assignmentRepo struct{}

type AssigmentRepository interface {
	AssignBotToDelivery(assignment *entity.Assignment) (*entity.Assignment, error)
}

func NewAssignmentRepository() AssigmentRepository {
	return &assignmentRepo{}
}

func (*assignmentRepo) AssignBotToDelivery(assignment *entity.Assignment) (*entity.Assignment, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./serviceAccount.json")
	client, err := firestore.NewClient(ctx, projectId, sa)
	if err != nil {
		log.Fatalf("Failed to create a Firestore app: %v", err)
		return nil, err
	}
	defer client.Close()

	delivery, err := NewDeliveryRepository().GetDocumentById(assignment.DeliveryId)
	if err != nil {
		log.Fatalf("Failed to retrieve a delivery document: %v", err)
		return nil, err
	}
	if delivery["state"] != "pending" {
		log.Println("Error: Delivery state is not pending")
		return nil, errors.New("delivery is not pending")
	}

	bot, err := NewBotRepository().GetBotbyId(assignment.BotId)
	if err != nil {
		log.Fatalf("Failed to retrieve a bot document: %v", err)
		return nil, err
	}
	if bot["status"] != "available" {
		log.Println("Error: Bot status is not available")
		return nil, errors.New("Bot status is not available")
	}

	_, _, err = client.Collection(assignmentsCollection).Add(ctx, assignment)
	if err != nil {
		log.Fatalf("Failed adding a new assignment: %v", err)
		return nil, err
	}

	return assignment, nil
}

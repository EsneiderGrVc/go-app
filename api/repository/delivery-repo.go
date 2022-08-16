package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/EsneiderGrVc/go_server/entity"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type deliveryRepo struct{}

type DeliveryRepository interface {
	Save(delivery *entity.Delivery) (*entity.Delivery, error)
	GetAll() ([]map[string]interface{}, error)
	GetDocumentById(id string) (map[string]interface{}, error)
}

// New Firestore Repository creates a new repo
func NewDeliveryRepository() DeliveryRepository {
	return &deliveryRepo{}
}

func (*deliveryRepo) Save(delivery *entity.Delivery) (*entity.Delivery, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./serviceAccount.json")
	client, err := firestore.NewClient(ctx, projectId, sa)
	if err != nil {
		log.Fatalf("Failed to create a Firestore app: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(deliveriesCollection).Add(ctx, map[string]interface{}{
		"id":            delivery.Id,
		"creation_date": delivery.CreationDate,
		"state":         delivery.State,
		"pickup": map[string]interface{}{
			"pickup_lat": delivery.Pickup.PickupLat,
			"pickup_lon": delivery.Pickup.PickupLon,
		},
		"dropoff": map[string]interface{}{
			"dropoff_lat": delivery.Dropoff.DropoffLat,
			"dropoff_lon": delivery.Dropoff.DropoffLon,
		},
		"zone_id": delivery.ZoneId,
	})
	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}
	return delivery, nil
}

func (*deliveryRepo) GetAll() ([]map[string]interface{}, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./serviceAccount.json")
	client, err := firestore.NewClient(ctx, projectId, sa)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []map[string]interface{}
	iter := client.Collection(deliveriesCollection).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return nil, err
		}

		var post map[string]interface{}
		doc.DataTo(&post)
		posts = append(posts, post)
	}
	return posts, nil
}

func (*deliveryRepo) GetDocumentById(id string) (map[string]interface{}, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./serviceAccount.json")
	client, err := firestore.NewClient(ctx, projectId, sa)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()

	collection := client.Collection(deliveriesCollection)
	docs, err1 := collection.Where("id", "==", id).Documents(ctx).GetAll()
	if err1 != nil {
		log.Fatalf("Failed to get a Document by id: %v", err)
	}

	result := map[string]interface{}{}
	for _, doc := range docs {
		doc.DataTo(&result)
	}
	return result, err
}

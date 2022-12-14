package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/EsneiderGrVc/go_server/entity"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type botRepo struct{}

type BotRepository interface {
	Save(bot *entity.Bot) (*entity.Bot, error)
	GetAll() ([]map[string]interface{}, error)
	GetBotbyId(id string) (map[string]interface{}, error)
}

func NewBotRepository() BotRepository {
	return &botRepo{}
}

func (*botRepo) Save(bot *entity.Bot) (*entity.Bot, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./serviceAccount.json")
	client, err := firestore.NewClient(ctx, projectId, sa)
	if err != nil {
		log.Fatalf("Failed to create a Firestore app: %v", err)
		return nil, err
	}
	defer client.Close()

	if len(bot.Id) == 0 {
		bot.Id = uuid.Must(uuid.NewRandom()).String()
	}

	_, _, err = client.Collection(botsCollection).Add(ctx, map[string]interface{}{
		"id":     bot.Id,
		"status": bot.Status,
		"location": map[string]interface{}{
			"lat": bot.Location.Lat,
			"lon": bot.Location.Lon,
		},
		"zone_id": bot.ZoneId,
	})
	if err != nil {
		log.Fatalf("Failed adding a new bot: %v", err)
		return nil, err
	}
	return bot, nil
}

func (*botRepo) GetAll() ([]map[string]interface{}, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./serviceAccount.json")
	client, err := firestore.NewClient(ctx, projectId, sa)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	var bots []map[string]interface{}
	iter := client.Collection(botsCollection).OrderBy("zone_id", firestore.Asc).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return nil, err
		}

		var bot map[string]interface{}
		doc.DataTo(&bot)
		bots = append(bots, bot)
	}
	return bots, nil
}

func (*botRepo) GetBotbyId(id string) (map[string]interface{}, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./serviceAccount.json")
	client, err := firestore.NewClient(ctx, projectId, sa)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()

	collection := client.Collection(botsCollection)
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

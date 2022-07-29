package dbutil

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/EsneiderGrVc/go_server/entity"
	"github.com/EsneiderGrVc/go_server/repository"
)

func FillDeliveries(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM ds.deliveries")
	if err != nil {
		panic("ds.deliveries table is empty")
	}

	deliveries, err := repository.NewDeliveryRepository().GetAll()
	if err != nil {
		log.Fatal(err)
	}
	if len(deliveries) > 0 {
		return nil
	}

	fmt.Printf("Sending dummy deliveries data to Firestore...\n")
	for rows.Next() {
		var delivery entity.Delivery
		err := rows.Scan(
			&delivery.Id,
			&delivery.CreationDate,
			&delivery.State,
			&delivery.Pickup.PickupLat,
			&delivery.Pickup.PickupLon,
			&delivery.Dropoff.DropoffLat,
			&delivery.Dropoff.DropoffLon,
			&delivery.ZoneId,
		)
		if err != nil {
			log.Fatalf("Error getting a row: %v\n", err)
		}
		_, err1 := repository.NewDeliveryRepository().Save(&delivery)
		if err1 != nil {
			log.Fatalf("Error Saving a delivery from local db: %v", err1)
		}
	}
	rows.Close()
	return err
}

func FillBots(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM bs.bots")
	if err != nil {
		panic("ds.deliveries table is empty")
	}

	bots, err := repository.NewBotRepository().GetAll()
	if err != nil {
		log.Fatal(err)
	}
	if len(bots) > 0 {
		return nil
	}

	fmt.Printf("Sending dummy bots data to Firestore...\n")
	for rows.Next() {
		var bot entity.Bot
		err := rows.Scan(
			&bot.Id,
			&bot.Status,
			&bot.Location.Lat,
			&bot.Location.Lon,
			&bot.ZoneId,
		)
		if err != nil {
			log.Fatalf("Error getting a row: %v\n", err)
		}
		_, err1 := repository.NewBotRepository().Save(&bot)
		if err1 != nil {
			log.Fatalf("Error Saving a delivery from local db: %v", err1)
		}
	}
	return err
}

func FillFirestoreDb(db *sql.DB) error {
	err := FillDeliveries(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ds.deliveries is loaded")
	err = FillBots(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("bs.bots is loaded")

	return err
}

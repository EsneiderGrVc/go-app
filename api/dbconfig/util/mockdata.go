package dbutil

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
)

func MockData(db *sql.DB) error {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	log.Println("Creating dummy deliveries data in postgres...")
	deliveriesQuery, err := ioutil.ReadFile(pwd + "/dbconfig/deliveries.sql")
	if err != nil {
		panic(err)
	}
	_, err2 := db.Exec(string(deliveriesQuery))
	if err2 != nil {
		log.Fatal(err2)
	}

	log.Println("Creating dummy bots data in postgres...")
	botsQuery, err := ioutil.ReadFile(pwd + "/dbconfig/bots.sql")
	if err != nil {
		panic(err)
	}
	_, err3 := db.Exec(string(botsQuery))
	if err3 != nil {
		log.Fatal(err3)
	}
	return err
}

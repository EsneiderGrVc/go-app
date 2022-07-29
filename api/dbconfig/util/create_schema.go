package dbutil

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
)

func CreateSchema(db *sql.DB) error {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	configQuery, err := ioutil.ReadFile(pwd + "/dbconfig/config-db.sql")
	if err != nil {
		panic(err)
	}

	_, err1 := db.Exec(string(configQuery))
	if err1 != nil {
		log.Fatal(err)
	}
	return err
}

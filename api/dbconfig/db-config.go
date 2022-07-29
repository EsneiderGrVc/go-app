package dbconfig

import (
	"database/sql"
	"log"

	dbutil "github.com/EsneiderGrVc/go_server/dbconfig/util"
)

type dbConfig struct {
}

type PGDB interface {
	Config()
}

func NewDbConfig() PGDB {
	return &dbConfig{}
}

func (c *dbConfig) Config() {
	connStr := "postgres://pguser:pguser@database:5432/kiwidb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	errSa := dbutil.CreateSchema(db)
	if errSa != nil {
		log.Fatal("Error Creting Schema: ", errSa)
	}

	rows, err := db.Query("SELECT * FROM ds.deliveries")
	if err != nil {
		panic("ds.deliveries table is empty")
	}

	if rows.Next() {
		errDData := dbutil.FillFirestoreDb(db)
		if errDData != nil {
			log.Fatal("Error Creting Data: ", errDData)
		}
		return
	} else {
		errDb := dbutil.MockData(db)
		if errDb != nil {
			log.Fatal("Error Creting Data: ", errDb)
		}
		c.Config()
	}
	defer rows.Close()
}

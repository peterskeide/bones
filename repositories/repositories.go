package repositories

import (
	"bones/config"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

var NotFoundError = sql.ErrNoRows

func Connect(dbconfig config.DatabaseConfig) {
	var err error
	db, err = sql.Open("postgres", dbconfig.ConnectionString())

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")
}

func Cleanup() {
	db.Close()
	log.Println("Closed database connection")
}

package repositories

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

var NotFoundError = sql.ErrNoRows

func Connect(conninfo string) {
	var err error
	db, err = sql.Open("postgres", conninfo)

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

type EntityFinder interface {
	Find(id int) (interface{}, error)
}

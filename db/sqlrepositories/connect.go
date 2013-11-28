package sqlrepositories

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/peterskeide/veil"
	"log"
)

var db *sql.DB
var dbveil veil.Veil

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

	dbveil = veil.New(db)

	log.Println("Connected to database")
}

func Cleanup() {
	err := dbveil.Close()

	if err != nil {
		log.Println(err)
	}

	err = db.Close()

	if err != nil {
		log.Println(err)
	}

	log.Println("Closed database connection")
}

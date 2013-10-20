package repositories

import (
	"labix.org/v2/mgo"
	"log"
)

var session *mgo.Session
var database string

var NotFoundError = mgo.ErrNotFound

func Connect(connInfo string, db string) {
	database = db

	var err error
	session, err = mgo.Dial(connInfo)

	if err != nil {
		panic(err)
	}

	log.Printf("Connected to MongoDB database '%s' on %s\n", db, connInfo)
}

func Cleanup() {
	session.Close()
	log.Println("Closed mgo session")
}

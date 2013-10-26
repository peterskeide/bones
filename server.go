package main

import (
	"bones/config"
	"bones/repositories"
	"bones/web/handlers"
	"github.com/gorilla/pat"
	"log"
	"net/http"
)

func main() {
	repositories.Connect(config.Database())
	defer repositories.Cleanup()

	r := pat.New()

	setupRouting(r)

	http.Handle("/", r)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	log.Println("Starting server on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupRouting(r *pat.Router) {
	r.Get("/signup", handlers.LoadSignupPage)
	r.Post("/signup", handlers.CreateNewUser)
	r.Get("/", handlers.LoadHomePage)
}

package main

import (
	"bones/config"
	"bones/repositories"
	"bones/web/filters"
	"bones/web/handlers"
	"github.com/gorilla/pat"
	"log"
	"net/http"
	"os"
)

func main() {
	repositories.Connect(config.DatabaseConnectionString())
	defer repositories.Cleanup()

	repositories.EnableSessions()

	r := setupRouting()

	http.Handle("/", r)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	port := portFromEnvOrDefault()
	log.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func setupRouting() *pat.Router {
	r := pat.New()

	handlers.SetRouter(r)

	r.Get("/users/{id:[0-9]+}/profile", filters.ApplyTo(handlers.LoadUserProfilePage, filters.Authenticate, filters.Params)).Name("userProfile")

	r.Get("/signup", handlers.LoadSignupPage)
	r.Post("/signup", handlers.CreateNewUser)

	r.Get("/login", handlers.LoadLoginPage)
	r.Post("/login", filters.ApplyTo(handlers.CreateNewSession, filters.Csrf))
	r.Get("/logout", handlers.Logout)

	r.Get("/", handlers.LoadHomePage)

	return r
}

func portFromEnvOrDefault() string {
	port := os.Getenv("PORT")

	if port != "" {
		return port
	}

	return "8080"
}

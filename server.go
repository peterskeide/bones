package main

import (
	"bones/config"
	"bones/repositories"
	"bones/web/filters"
	"bones/web/handlers"
	"bones/web/services"
	"bones/web/templating"
	"github.com/gorilla/pat"
	"log"
	"net/http"
	"os"
)

// Router
var r *pat.Router

// Services
var templateRenderer templating.TemplateRenderer
var shortcuts services.Shortcuts

// Filters
var f *filters.Filters

// Handlers
var homeHandler *handlers.HomeHandler
var loginHandler *handlers.LoginHandler
var signupHandler *handlers.SignupHandler

func main() {
	repositories.Connect(config.DatabaseConnectionString())
	defer repositories.Cleanup()

	repositories.EnableSessions()

	setupDependencies()

	setupRouting()

	http.Handle("/", r)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	port := portFromEnvOrDefault()
	log.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func setupDependencies() {
	r = pat.New()
	handlers.SetRouter(r)

	templateRenderer = templating.NewTemplateRenderer()
	shortcuts = &services.TemplatingShortcuts{templateRenderer}

	f = &filters.Filters{shortcuts}

	homeHandler = &handlers.HomeHandler{shortcuts}
	loginHandler = &handlers.LoginHandler{shortcuts}
	signupHandler = &handlers.SignupHandler{shortcuts}
}

func setupRouting() {
	r.Get("/users/{id:[0-9]+}/profile", filters.ApplyTo(loginHandler.LoadUserProfilePage, f.Authenticate, filters.Params)).Name("userProfile")

	r.Get("/signup", signupHandler.LoadSignupPage)
	r.Post("/signup", signupHandler.CreateNewUser)

	r.Get("/login", loginHandler.LoadLoginPage)
	r.Post("/login", filters.ApplyTo(loginHandler.CreateNewSession, filters.Csrf))
	r.Get("/logout", loginHandler.Logout)

	r.Get("/", homeHandler.LoadHomePage)
}

func portFromEnvOrDefault() string {
	port := os.Getenv("PORT")

	if port != "" {
		return port
	}

	return "8080"
}

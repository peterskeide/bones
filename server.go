package main

import (
	"bones/config"
	"bones/db/sqlrepositories"
	"bones/repositories"
	"bones/web/authentication"
	"bones/web/filters"
	"bones/web/handlers"
	"bones/web/sessions"
	"bones/web/templating"
	"github.com/gorilla/pat"
	"log"
	"net/http"
	"os"
)

// Router
var r *pat.Router

// Services
var templateRenderer handlers.TemplateRenderer
var shortcuts handlers.Shortcuts
var authenticator handlers.Authenticator
var sessionStore handlers.SessionStore

// Repositories
var userRepository repositories.UserRepository

// Filters
var f *filters.Filters

// Handlers
var homeHandler *handlers.HomeHandler
var loginHandler *handlers.LoginHandler
var signupHandler *handlers.SignupHandler

func main() {
	sqlrepositories.Connect(config.DatabaseConnectionString())
	defer sqlrepositories.Cleanup()

	sessions.Enable()

	setupDependencies()

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/favicon.ico", func(req http.ResponseWriter, res *http.Request) {
		http.ServeFile(req, res, "./assets/images/favicon.png")
	})

	http.HandleFunc("/robots.txt", func(req http.ResponseWriter, res *http.Request) {
		http.ServeFile(req, res, "./assets/robots.txt")
	})

	setupRouting()

	http.Handle("/", r)

	port := portFromEnvOrDefault()
	log.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func setupDependencies() {
	r = pat.New()
	handlers.SetRouter(r)

	userRepository := sqlrepositories.NewUserRepository()

	authenticator = &authentication.EmailAuthenticator{userRepository}
	sessionStore = &sessions.CookieSessionStore{}

	templateRenderer = templating.NewTemplateRenderer()
	shortcuts = handlers.Shortcuts{templateRenderer, sessionStore}

	f = &filters.Filters{shortcuts, sessionStore, userRepository}

	homeHandler = &handlers.HomeHandler{shortcuts, userRepository}
	loginHandler = &handlers.LoginHandler{shortcuts, authenticator, userRepository, sessionStore}
	signupHandler = &handlers.SignupHandler{shortcuts, userRepository}
}

func setupRouting() {
	r.Get("/users/{id:[0-9]+}/profile", filters.ApplyTo(loginHandler.LoadUserProfilePage, f.Authenticate, filters.Params)).Name("userProfile")

	r.Get("/signup", signupHandler.LoadSignupPage)
	r.Post("/signup", signupHandler.CreateNewUser)

	r.Get("/login", loginHandler.LoadLoginPage)
	r.Post("/login", filters.ApplyTo(loginHandler.CreateNewSession, f.Csrf))
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

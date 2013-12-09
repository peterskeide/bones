package sessions

import (
	"bones/web/handlers"
	"encoding/hex"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"os"
)

var sessionStore sessions.Store

// Initialize the session store with authentication
// and encryption keys.
//
// Keys are read from the environment variables
// SESSION_AUTH_KEY and SESSION_ENCRYPTION_KEY
// If not set, temporary keys will be used.
//
// Using temporary keys will invalidate all sessions
// when the server restarts.
func Enable() {
	auth_key := []byte(os.Getenv("SESSION_AUTH_KEY"))

	if len(auth_key) == 0 {
		log.Println("Using temporary authentication key for session")
		auth_key = securecookie.GenerateRandomKey(64)
	}

	encryption_key := []byte(os.Getenv("SESSION_ENCRYPTION_KEY"))

	if len(encryption_key) == 0 {
		log.Println("Using temporary encryption key for session")
		encryption_key = securecookie.GenerateRandomKey(32)
	}

	sessionStore = sessions.NewCookieStore(auth_key, encryption_key)
}

type CookieSessionStore struct{}

func (store CookieSessionStore) Session(res http.ResponseWriter, req *http.Request) handlers.Session {
	session, _ := sessionStore.Get(req, "bones_session")
	return &CookieSession{session, res, req}
}

type CookieSession struct {
	session        *sessions.Session
	responseWriter http.ResponseWriter
	request        *http.Request
}

func (s *CookieSession) UserId() int {
	value := s.session.Values["userId"]

	if id, ok := value.(int); ok {
		return id
	}

	return -1
}

func (s *CookieSession) SetUserId(id int) {
	s.session.Values["userId"] = id
}

// Returns existing CsrfToken from cookie or
// creates and saves a new token.
// Will panic if it is unable to return a token.
func (s *CookieSession) CsrfToken() string {
	token, ok := s.session.Values["CsrfToken"].(string)

	if !ok {
		randomKey := securecookie.GenerateRandomKey(32)
		token = hex.EncodeToString(randomKey)
		s.session.Values["CsrfToken"] = token

		err := s.Save()

		if err != nil {
			panic(err)
		}
	}

	return token
}

func (s *CookieSession) Clear() error {
	s.session.Values = nil
	return s.Save()
}

func (s *CookieSession) Save() error {
	return s.session.Save(s.request, s.responseWriter)
}

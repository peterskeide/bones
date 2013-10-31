package repositories

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
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
func EnableSessions() {
	auth_key := []byte(os.Getenv("SESSION_AUTH_KEY"))

	if len(auth_key) == 0 {
		auth_key = securecookie.GenerateRandomKey(64)
	}

	encryption_key := []byte(os.Getenv("SESSION_ENCRYPTION_KEY"))

	if len(encryption_key) == 0 {
		encryption_key = securecookie.GenerateRandomKey(32)
	}

	sessionStore = sessions.NewCookieStore(auth_key, encryption_key)
}

type SessionRepository interface {
	Value(key string) interface{}
	SetValue(key string, value interface{})
	Save() error
}

func Session(res http.ResponseWriter, req *http.Request) SessionRepository {
	session, _ := sessionStore.Get(req, "bones_session")
	return &CookieSessionRepository{session, res, req}
}

type CookieSessionRepository struct {
	session        *sessions.Session
	responseWriter http.ResponseWriter
	request        *http.Request
}

func (s *CookieSessionRepository) Value(key string) interface{} {
	return s.session.Values[key]
}

func (s *CookieSessionRepository) SetValue(key string, value interface{}) {
	s.session.Values[key] = value
}

func (s *CookieSessionRepository) Save() error {
	return s.session.Save(s.request, s.responseWriter)
}

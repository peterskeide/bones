package context

import (
	"bones/entities"
	"github.com/gorilla/context"
	"net/http"
)

type key int

const (
	currentUser key = 0
	params          = 1
)

func SetCurrentUser(req *http.Request, user *entities.User) {
	context.Set(req, currentUser, user)
}

func CurrentUser(req *http.Request) *entities.User {
	if user, ok := context.Get(req, currentUser).(*entities.User); ok {
		return user
	}

	return nil
}

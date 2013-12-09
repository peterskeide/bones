package filters

import (
	"bones/repositories"
	"bones/web/context"
	"bones/web/handlers"
	"log"
	"net/http"
)

type Filters struct {
	handlers.Shortcuts
	SessionStore handlers.SessionStore
	Users        repositories.UserRepository
}

func (f *Filters) Authenticate(res http.ResponseWriter, req *http.Request, chain *RequestFilterChain) {
	session := f.SessionStore.Session(res, req)
	id := session.UserId()
	user, err := f.Users.FindById(id)

	if err != nil {
		if err != repositories.NotFoundError {
			log.Println("Error when finding user for authentication:", err)
		}

		f.RedirectToLogin(res, req)

		return
	}

	context.SetCurrentUser(req, user)
	chain.next()
}

func (f *Filters) Csrf(res http.ResponseWriter, req *http.Request, chain *RequestFilterChain) {
	session := f.SessionStore.Session(res, req)
	sessionToken := session.CsrfToken()
	formToken := req.FormValue("CsrfToken")

	if sessionToken != formToken {
		http.Error(res, "Forbidden", http.StatusForbidden)

		return
	}

	chain.next()
}

func Params(res http.ResponseWriter, req *http.Request, chain *RequestFilterChain) {
	err := context.InitParams(req)

	if err != nil {
		http.Error(res, "Bad request", http.StatusBadRequest)

		return
	}

	chain.next()
}

package filters

import (
	"bones/repositories"
	"bones/web/context"
	"log"
	"net/http"
)

func Params(res http.ResponseWriter, req *http.Request, chain *RequestFilterChain) {
	err := context.InitParams(req)

	if err != nil {
		http.Error(res, "Bad request", http.StatusBadRequest)

		return
	}

	chain.next()
}

func Authenticate(res http.ResponseWriter, req *http.Request, chain *RequestFilterChain) {
	session := repositories.Session(res, req)
	value := session.Value("user_id")

	if id, ok := value.(int); ok {
		user, err := repositories.Users.FindById(id)

		if err != nil {
			if err != repositories.NotFoundError {
				log.Println("Error when finding user for authentication:", err)
			}

			redirectToLogin(res, req)

			return
		}

		context.SetCurrentUser(req, user)
		chain.next()
	} else {
		redirectToLogin(res, req)
	}
}

func redirectToLogin(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/login", 302)
}

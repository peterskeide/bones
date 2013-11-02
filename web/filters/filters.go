package filters

import (
	"bones/repositories"
	"bones/web/actions"
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

			actions.RedirectToLogin(res, req)

			return
		}

		context.SetCurrentUser(req, user)
		chain.next()
	} else {
		actions.RedirectToLogin(res, req)
	}
}

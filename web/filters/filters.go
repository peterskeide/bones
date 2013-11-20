package filters

import (
	"bones/repositories"
	"bones/web/context"
	"bones/web/services"
	"log"
	"net/http"
)

type Filters struct {
	services.Shortcuts
}

func (f *Filters) Authenticate(res http.ResponseWriter, req *http.Request, chain *RequestFilterChain) {
	session := repositories.Session(res, req)
	value := session.Value("user_id")

	if id, ok := value.(int); ok {
		user, err := repositories.Users().FindById(id)

		if err != nil {
			if err != repositories.NotFoundError {
				log.Println("Error when finding user for authentication:", err)
			}

			f.RedirectToLogin(res, req)

			return
		}

		context.SetCurrentUser(req, user)
		chain.next()
	} else {
		f.RedirectToLogin(res, req)
	}
}

func Params(res http.ResponseWriter, req *http.Request, chain *RequestFilterChain) {
	err := context.InitParams(req)

	if err != nil {
		http.Error(res, "Bad request", http.StatusBadRequest)

		return
	}

	chain.next()
}

func Csrf(res http.ResponseWriter, req *http.Request, chain *RequestFilterChain) {
	session := repositories.Session(res, req)
	sessionToken, ok := session.Value("CsrfToken").(string)

	if ok {
		formToken := req.FormValue("CsrfToken")

		if sessionToken == "" || sessionToken != formToken {
			http.Error(res, "Forbidden", http.StatusForbidden)

			return
		}

		chain.next()
	} else {
		http.Error(res, "Forbidden", http.StatusForbidden)
	}
}

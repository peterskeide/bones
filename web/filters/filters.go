package filters

import (
	"bones/repositories"
	"net/http"
)

func Authenticate(res http.ResponseWriter, req *http.Request, chain *RequestFilterChain) {
	session := repositories.Session(res, req)
	value := session.Value("user_id")

	if _, ok := value.(int); ok {
		chain.next()

		return
	}

	http.Redirect(res, req, "/login", 302)
}

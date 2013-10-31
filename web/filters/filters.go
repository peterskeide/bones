package filters

import (
	"bones/repositories"
	"net/http"
)

func Authenticate(res http.ResponseWriter, req *http.Request) bool {
	session := repositories.Session(res, req)
	value := session.Value("user_id")

	if _, ok := value.(int); ok {
		return true
	}

	http.Redirect(res, req, "/login", 302)
	return false
}

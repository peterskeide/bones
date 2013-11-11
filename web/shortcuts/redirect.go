package shortcuts

import (
	"net/http"
)

func RedirectToLogin(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/login", http.StatusFound)
}

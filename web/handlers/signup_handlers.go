package handlers

import (
	"bones/repositories"
	"bones/web/forms"
	"bones/web/services"
	"bones/web/templating"
	"net/http"
)

type SignupHandler struct {
	services.Shortcuts
	Users repositories.UserRepository
}

func (h *SignupHandler) LoadSignupPage(res http.ResponseWriter, req *http.Request) {
	h.RenderPage(res, newSignupContext())
}

func (h *SignupHandler) CreateNewUser(res http.ResponseWriter, req *http.Request) {
	form := forms.SignupForm{Users: h.Users}
	err := h.ProcessForm(req, &form)

	if err != nil {
		h.RenderPageWithErrors(res, newSignupContext(), err)

		return
	}

	http.Redirect(res, req, "/", http.StatusFound)
}

func newSignupContext() *templating.BaseContext {
	return templating.NewBaseContext("signup.html")
}

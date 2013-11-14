package handlers

import (
	"bones/web/forms"
	"bones/web/services"
	"bones/web/templating"
	"net/http"
)

type SignupHandler struct {
	services.Shortcuts
}

func (h *SignupHandler) LoadSignupPage(res http.ResponseWriter, req *http.Request) {
	h.RenderPage(res, newSignupContext())
}

func (h *SignupHandler) CreateNewUser(res http.ResponseWriter, req *http.Request) {
	err := h.ProcessForm(req, new(forms.SignupForm))

	if err != nil {
		h.RenderPageWithErrors(res, newSignupContext(), err)

		return
	}

	http.Redirect(res, req, "/", http.StatusFound)
}

func newSignupContext() *templating.BaseContext {
	return templating.NewBaseContext("signup.html")
}

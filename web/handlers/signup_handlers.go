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
	h.RenderPage(res, h.newSignupContext(res, req))
}

func (h *SignupHandler) CreateNewUser(res http.ResponseWriter, req *http.Request) {
	err := h.validateInputAndCreateUser(req)

	if err != nil {
		h.RenderPageWithErrors(res, h.newSignupContext(res, req), err)
	} else {
		http.Redirect(res, req, "/", http.StatusFound)
	}
}

func (h *SignupHandler) validateInputAndCreateUser(req *http.Request) error {
	form := forms.SignupForm{}

	err := h.DecodeAndValidate(req, &form)

	if err != nil {
		return err
	}

	user, err := form.User()

	if err != nil {
		return err
	}

	return h.Users.Insert(user)
}

func (h *SignupHandler) newSignupContext(res http.ResponseWriter, req *http.Request) *templating.BaseContext {
	return h.TemplateContext(res, req, "signup.html")
}

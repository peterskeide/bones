package handlers

import (
	"bones/repositories"
	"bones/web/forms"
	"net/http"
)

type SignupHandler struct {
	Shortcuts
	Users repositories.UserRepository
}

func (h *SignupHandler) LoadSignupPage(res http.ResponseWriter, req *http.Request) {
	h.RenderPage(res, req, h.newSignupContext(res, req))
}

func (h *SignupHandler) CreateNewUser(res http.ResponseWriter, req *http.Request) {
	err := h.validateInputAndCreateUser(req)

	if err != nil {
		h.RenderPageWithErrors(res, req, h.newSignupContext(res, req), err)
	} else {
		h.AddFlashNotice(res, req, "User created")
		h.redirect(res, req, "/")
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

func (h *SignupHandler) newSignupContext(res http.ResponseWriter, req *http.Request) *BaseContext {
	return h.TemplateContext(res, req, "signup.html")
}

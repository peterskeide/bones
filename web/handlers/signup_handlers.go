package handlers

import (
	"bones/entities"
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
	err := h.validateInputAndCreateUser(req)

	if err != nil {
		h.RenderPageWithErrors(res, newSignupContext(), err)
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

	encryptedPassword, err := form.EncryptedPassword()

	if err != nil {
		return err
	}

	user := entities.User{
		Email:    form.Email,
		Password: encryptedPassword,
	}

	return h.Users.Insert(&user)
}

func newSignupContext() *templating.BaseContext {
	return templating.NewBaseContext("signup.html")
}

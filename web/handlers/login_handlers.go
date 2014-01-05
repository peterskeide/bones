package handlers

import (
	"bones/entities"
	"bones/repositories"
	"bones/web/context"
	"bones/web/forms"
	"log"
	"net/http"
	"strconv"
)

type Authenticator interface {
	Authenticate(login string, password string) (*entities.User, error)
}

type ProfileContext struct {
	*BaseContext
	User *entities.User
}

type LoginHandler struct {
	Shortcuts
	Authenticator
	Users        repositories.UserRepository
	SessionStore SessionStore
}

func (h *LoginHandler) LoadLoginPage(res http.ResponseWriter, req *http.Request) {
	ctx := h.TemplateContext(res, req, "login.html")
	h.RenderPage(res, req, ctx)
}

func (h *LoginHandler) CreateNewSession(res http.ResponseWriter, req *http.Request) {
	user, err := h.validateCredentialsAndLoginUser(res, req)

	if err != nil {
		ctx := h.TemplateContext(res, req, "login.html")
		h.RenderPageWithErrors(res, req, ctx, err)
	} else {
		h.AddFlashNotice(res, req, "Login successful")
		url := routeURL("userProfile", "id", strconv.Itoa(user.Id))
		h.redirect(res, req, url)
	}
}

func (h *LoginHandler) validateCredentialsAndLoginUser(res http.ResponseWriter, req *http.Request) (*entities.User, error) {
	form := forms.LoginForm{}

	err := h.DecodeAndValidate(req, &form)

	if err != nil {
		return nil, err
	}

	user, err := h.Authenticate(form.Email, form.Password)

	if err != nil {
		return nil, err
	}

	session := h.SessionStore.Session(res, req)
	session.SetUserId(user.Id)

	return user, nil
}

func (h *LoginHandler) LoadUserProfilePage(res http.ResponseWriter, req *http.Request) {
	id, _ := context.Params(req).GetInt(":id")

	// A user can only see his/her own profile
	if context.CurrentUser(req).Id != id {
		h.Render401(res, req)

		return
	}

	entity := h.FindEntityOr404(res, req, h.Users, id)

	if user, ok := entity.(*entities.User); ok {
		ctx := ProfileContext{h.TemplateContext(res, req, "profile.html"), user}
		h.RenderPage(res, req, &ctx)
	}
}

func (h *LoginHandler) Logout(res http.ResponseWriter, req *http.Request) {
	session := h.SessionStore.Session(res, req)
	err := session.Clear()

	if err != nil {
		log.Println("Error when clearing session:", err)
	}

	h.RedirectToLogin(res, req)
}

package handlers

import (
	"bones/entities"
	"bones/repositories"
	"bones/web/context"
	"bones/web/forms"
	"bones/web/services"
	"bones/web/templating"
	"log"
	"net/http"
	"strconv"
)

type ProfileContext struct {
	*templating.BaseContext
	User *entities.User
}

type LoginHandler struct {
	services.Shortcuts
}

func (h *LoginHandler) LoadLoginPage(res http.ResponseWriter, req *http.Request) {
	if ctx := h.FormContextOr500(res, req, "login.html"); ctx != nil {
		h.RenderPage(res, ctx)
	}
}

func (h *LoginHandler) CreateNewSession(res http.ResponseWriter, req *http.Request) {
	form := forms.LoginForm{ResponseWriter: res, Request: req}
	err := h.ProcessForm(req, &form)

	if err != nil {
		if ctx := h.FormContextOr500(res, req, "login.html"); ctx != nil {
			h.RenderPageWithErrors(res, ctx, err)
		}

		return
	}

	url := routeURL("userProfile", "id", strconv.Itoa(form.User.Id))
	http.Redirect(res, req, url, http.StatusFound)
}

func (h *LoginHandler) LoadUserProfilePage(res http.ResponseWriter, req *http.Request) {
	id, _ := context.Params(req).GetInt(":id")

	// A user can only see his/her own profile
	if context.CurrentUser(req).Id != id {
		h.Render401(res)

		return
	}

	entity := h.FindEntityOr404(res, req, repositories.Users, id)

	if user, ok := entity.(*entities.User); ok {
		ctx := ProfileContext{templating.NewBaseContext("profile.html"), user}
		h.RenderPage(res, &ctx)
	}
}

func (h *LoginHandler) Logout(res http.ResponseWriter, req *http.Request) {
	session := repositories.Session(res, req)
	err := session.Clear()

	if err != nil {
		log.Println("Error when clearing session:", err)
	}

	h.RedirectToLogin(res, req)
}

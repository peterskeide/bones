package handlers

import (
	"bones/entities"
	"bones/repositories"
	"bones/web/actions"
	"bones/web/context"
	"bones/web/forms"
	"bones/web/templating"
	"log"
	"net/http"
	"strconv"
)

type ProfileContext struct {
	*templating.BaseContext
	User *entities.User
}

func LoadLoginPage(res http.ResponseWriter, req *http.Request) {
	actions.RenderPage(res, newLoginContext())
}

func CreateNewSession(res http.ResponseWriter, req *http.Request) {
	form := forms.LoginForm{ResponseWriter: res, Request: req}
	err := actions.ProcessForm(req, &form)

	if err != nil {
		actions.RenderPageWithErrors(res, newLoginContext(), err)

		return
	}

	url := routeURL("userProfile", "id", strconv.Itoa(form.User.Id))
	http.Redirect(res, req, url, 302)
}

func LoadUserProfilePage(res http.ResponseWriter, req *http.Request) {
	// If GetInt returns an error, id will be 0 and the entity lookup will fail.
	// Consider handling the error to avoid unnecessary database requests.
	id, _ := context.Params(req).GetInt(":id")
	entity := actions.FindEntityOr404(res, req, repositories.Users, id)

	if user, ok := entity.(*entities.User); ok {
		ctx := ProfileContext{templating.NewBaseContext("profile.html"), user}
		actions.RenderPage(res, &ctx)
	}
}

func Logout(res http.ResponseWriter, req *http.Request) {
	session := repositories.Session(res, req)
	err := session.Clear()

	if err != nil {
		log.Println("Error when clearing session:", err)
	}

	actions.RedirectToLogin(res, req)
}

func newLoginContext() *templating.BaseContext {
	return templating.NewBaseContext("login.html")
}

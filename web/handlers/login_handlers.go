package handlers

import (
	"bones/entities"
	"bones/repositories"
	"bones/web/shortcuts"
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
	if ctx := shortcuts.FormContextOr500(res, req, "login.html"); ctx != nil {
		shortcuts.RenderPage(res, ctx)
	}
}

func CreateNewSession(res http.ResponseWriter, req *http.Request) {
	form := forms.LoginForm{ResponseWriter: res, Request: req}
	err := shortcuts.ProcessForm(req, &form)

	if err != nil {
		if ctx := shortcuts.FormContextOr500(res, req, "login.html"); ctx != nil {
			shortcuts.RenderPageWithErrors(res, ctx, err)
		}

		return
	}

	url := routeURL("userProfile", "id", strconv.Itoa(form.User.Id))
	http.Redirect(res, req, url, http.StatusFound)
}

func LoadUserProfilePage(res http.ResponseWriter, req *http.Request) {
	// If GetInt returns an error, id will be 0 and the entity lookup will fail.
	// Consider handling the error to avoid unnecessary database requests.
	id, _ := context.Params(req).GetInt(":id")

	// A user can only se his/her own profile
	if context.CurrentUser(req).Id != id {
		shortcuts.Render401(res)

		return
	}

	entity := shortcuts.FindEntityOr404(res, req, repositories.Users, id)

	if user, ok := entity.(*entities.User); ok {
		ctx := ProfileContext{templating.NewBaseContext("profile.html"), user}
		shortcuts.RenderPage(res, &ctx)
	}
}

func Logout(res http.ResponseWriter, req *http.Request) {
	session := repositories.Session(res, req)
	err := session.Clear()

	if err != nil {
		log.Println("Error when clearing session:", err)
	}

	shortcuts.RedirectToLogin(res, req)
}

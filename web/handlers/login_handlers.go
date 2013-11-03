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
	ctx, err := newLoginContext(res, req)

	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)

		return
	}

	actions.RenderPage(res, ctx)
}

func CreateNewSession(res http.ResponseWriter, req *http.Request) {
	form := forms.LoginForm{ResponseWriter: res, Request: req}
	err := actions.ProcessForm(req, &form)

	if err != nil {
		ctx, err := newLoginContext(res, req)

		if err != nil {
			http.Error(res, "Internal server error", http.StatusInternalServerError)

			return
		}

		actions.RenderPageWithErrors(res, ctx, err)

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
		actions.Render401(res)

		return
	}

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

func newLoginContext(res http.ResponseWriter, req *http.Request) (*templating.FormContext, error) {
	base := templating.NewBaseContext("login.html")
	return templating.NewFormContext(res, req, base)
}

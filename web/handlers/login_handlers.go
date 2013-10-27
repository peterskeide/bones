package handlers

import (
	"bones/web/actions"
	"bones/web/forms"
	//"fmt"
	"net/http"
)

func LoadLoginPage(res http.ResponseWriter, req *http.Request) {
	actions.RenderPage(res, newLoginContext())
}

func CreateNewSession(res http.ResponseWriter, req *http.Request) {
	form := &forms.LoginForm{Request: req}
	err := actions.ProcessForm(req, form)

	if err != nil {
		ctx := newLoginContext()
		ctx.AddError(err)
		actions.RenderPage(res, ctx)

		return
	}

	// url := fmt.Sprintf("/users/%d/profile", form.User.Id)
	http.Redirect(res, req, "/", 302)
}

func newLoginContext() *BaseContext {
	return newBaseContext("login.html")
}

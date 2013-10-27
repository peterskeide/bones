package handlers

import (
	"bones/web/actions"
	"bones/web/forms"
	"net/http"
)

func LoadSignupPage(res http.ResponseWriter, req *http.Request) {
	actions.RenderPage(res, newSignupContext())
}

func CreateNewUser(res http.ResponseWriter, req *http.Request) {
	err := actions.ProcessForm(req, new(forms.SignupForm))

	if err != nil {
		ctx := newSignupContext()
		ctx.AddError(err)
		actions.RenderPage(res, ctx)

		return
	}

	http.Redirect(res, req, "/", 302)
}

func newSignupContext() *BaseContext {
	return newBaseContext("signup.html")
}

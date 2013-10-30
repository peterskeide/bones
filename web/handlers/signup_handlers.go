package handlers

import (
	"bones/web/actions"
	"bones/web/forms"
	"bones/web/templating"
	"net/http"
)

func LoadSignupPage(res http.ResponseWriter, req *http.Request) {
	actions.RenderPage(res, newSignupContext())
}

func CreateNewUser(res http.ResponseWriter, req *http.Request) {
	err := actions.ProcessForm(req, new(forms.SignupForm))

	if err != nil {
		actions.RenderPageWithErrors(res, newSignupContext(), err)

		return
	}

	http.Redirect(res, req, "/", 302)
}

func newSignupContext() *templating.BaseContext {
	return templating.NewBaseContext("signup.html")
}

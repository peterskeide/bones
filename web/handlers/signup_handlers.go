package handlers

import (
	"bones/web/shortcuts"
	"bones/web/forms"
	"bones/web/templating"
	"net/http"
)

func LoadSignupPage(res http.ResponseWriter, req *http.Request) {
	shortcuts.RenderPage(res, newSignupContext())
}

func CreateNewUser(res http.ResponseWriter, req *http.Request) {
	err := shortcuts.ProcessForm(req, new(forms.SignupForm))

	if err != nil {
		shortcuts.RenderPageWithErrors(res, newSignupContext(), err)

		return
	}

	http.Redirect(res, req, "/", http.StatusFound)
}

func newSignupContext() *templating.BaseContext {
	return templating.NewBaseContext("signup.html")
}

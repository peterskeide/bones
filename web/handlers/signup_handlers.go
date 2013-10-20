package handlers

import (
	"bones/web/actions"
	"bones/web/forms"
	"net/http"
)

func LoadSignupPage(res http.ResponseWriter, req *http.Request) {
	(actions.RenderPage{
		ResponseWriter: res,
		PageContext:    newSignupContext()}).Run()
}

func CreateNewUser(res http.ResponseWriter, req *http.Request) {
	(actions.SaveFormAndRedirect{
		ResponseWriter: res,
		Request:        req,
		Form:           new(forms.SignupForm),
		SuccessUrl:     "/",
		ErrorContext:   newSignupContext()}).Run()
}

func newSignupContext() *BaseContext {
	return newBaseContext("signup.html")
}

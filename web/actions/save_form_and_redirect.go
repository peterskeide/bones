package actions

import (
	"bones/web/forms"
	"bones/web/templating"
	"net/http"
)

type SaveFormAndRedirect struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Form           forms.Form
	SuccessUrl     string
	ErrorContext   templating.TemplateContext
}

func (wf SaveFormAndRedirect) Run() {
	if err := forms.DecodeForm(wf.Form, wf.Request); wf.renderError(err) {
		return
	}

	if err := wf.Form.Validate(); wf.renderError(err) {
		return
	}

	if err := wf.Form.Save(); wf.renderError(err) {
		return
	}

	http.Redirect(wf.ResponseWriter, wf.Request, wf.SuccessUrl, 302)
}

func (wf SaveFormAndRedirect) renderError(err error) bool {
	if err != nil {
		wf.ErrorContext.AddError(err)

		renderErr := templating.RenderTemplate(wf.ResponseWriter, wf.ErrorContext)

		if renderErr != nil {
			logTemplateRenderingErrorAndRespond500(wf.ResponseWriter, renderErr, wf.ErrorContext)
		}

		return true
	}

	return false
}

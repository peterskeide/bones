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
	forms.DecodeForm(wf.Form, wf.Request)

	err := wf.Form.Validate()

	if err != nil {
		wf.renderError(err)
	} else {
		err = wf.Form.Save()

		if err != nil {
			wf.renderError(err)
			return
		}

		http.Redirect(wf.ResponseWriter, wf.Request, wf.SuccessUrl, 302)
	}
}

func (wf SaveFormAndRedirect) renderError(err error) {
	wf.ErrorContext.AddError(err)
	renderErr := templating.RenderTemplate(wf.ResponseWriter, wf.ErrorContext)

	if renderErr != nil {
		logTemplateRenderingErrorAndRespond500(wf.ResponseWriter, renderErr, wf.ErrorContext)
	}
}

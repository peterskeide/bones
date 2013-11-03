package actions

import (
	"bones/web/forms"
	"bones/web/templating"
	"net/http"
)

func ProcessForm(req *http.Request, form forms.Form) error {
	if err := forms.DecodeForm(form, req); err != nil {
		return err
	}

	if err := form.Validate(); err != nil {
		return err
	}

	return form.Save()
}

func FormContextOr500(res http.ResponseWriter, req *http.Request, templateName string) *templating.FormContext {
	baseCtx := templating.NewBaseContext(templateName)
	formCtx, err := templating.NewFormContext(res, req, baseCtx)

	if err != nil {
		Render500(res, req)

		return nil
	}

	return formCtx
}

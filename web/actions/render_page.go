package actions

import (
	"bones/web/templating"
	"net/http"
)

type RenderPage struct {
	ResponseWriter http.ResponseWriter
	PageContext    templating.TemplateContext
}

func (wf RenderPage) Run() {
	err := templating.RenderTemplate(wf.ResponseWriter, wf.PageContext)

	if err != nil {
		logTemplateRenderingErrorAndRespond500(wf.ResponseWriter, err, wf.PageContext)
	}
}

package actions

import (
	"bones/web/templating"
	"net/http"
)

func RenderPageWithErrors(res http.ResponseWriter, pageContext templating.TemplateContext, errors ...error) {
	for _, err := range errors {
		pageContext.AddError(err)
	}

	RenderPage(res, pageContext)
}

package actions

import (
	"bones/web/templating"
	"log"
	"net/http"
)

func RenderPage(res http.ResponseWriter, pageContext templating.TemplateContext) {
	err := templating.RenderTemplate(res, pageContext)

	if err != nil {
		log.Println("Error when rendering template:", err, ". Context:", pageContext)
		http.Error(res, "Server encountered an error", 500)
	}
}

func RenderPageWithErrors(res http.ResponseWriter, pageContext templating.TemplateContext, errors ...error) {
	for _, err := range errors {
		pageContext.AddError(err)
	}

	RenderPage(res, pageContext)
}

func Render404(res http.ResponseWriter, req *http.Request) {
	err := templating.RenderTemplate(res, templating.NewBaseContext("404.html"))

	if err != nil {
		log.Println("Error when rendering 404 template:", err)
		http.NotFound(res, req)
	}
}

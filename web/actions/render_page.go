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

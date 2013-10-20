package actions

import (
	"bones/web/templating"
	"log"
	"net/http"
)

func logTemplateRenderingErrorAndRespond500(res http.ResponseWriter, err error, ctx templating.TemplateContext) {
	log.Println("Error when rendering template:", err, ". Context:", ctx)
	http.Error(res, "Server encountered an error", 500)
}

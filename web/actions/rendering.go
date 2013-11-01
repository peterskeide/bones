package actions

import (
	"bones/repositories"
	"bones/web/templating"
	"log"
	"net/http"
)

const (
	serverError string = "Server encountered an error"
)

func RenderPage(res http.ResponseWriter, pageContext templating.TemplateContext) {
	err := templating.RenderTemplate(res, pageContext)

	if err != nil {
		log.Println("Error when rendering template:", err, ". Context:", pageContext)
		http.Error(res, serverError, 500)
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

func FindEntityOr404(res http.ResponseWriter, req *http.Request, ef repositories.EntityFinder, id int) interface{} {
	entity, err := ef.Find(id)

	if err != nil {
		if err == repositories.NotFoundError {
			Render404(res, req)
		} else {
			log.Println("Error on entity find:", err)
			http.Error(res, serverError, 500)
		}

		return nil
	}

	return entity
}

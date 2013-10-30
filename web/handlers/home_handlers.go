package handlers

import (
	"bones/entities"
	"bones/repositories"
	"bones/web/actions"
	"bones/web/templating"
	"errors"
	"log"
	"net/http"
)

type HomeContext struct {
	*templating.BaseContext
	Users []entities.User
}

func LoadHomePage(res http.ResponseWriter, req *http.Request) {
	ctx := HomeContext{templating.NewBaseContext("index.html"), nil}
	users, err := repositories.Users.All()

	if err != nil {
		log.Println("Error loading users from repository:", err)
		ctx.AddError(errors.New("Unable to load users"))
	} else {
		ctx.Users = users
	}

	actions.RenderPage(res, ctx)
}

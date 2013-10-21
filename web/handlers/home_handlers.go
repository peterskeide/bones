package handlers

import (
	"bones/entities"
	"bones/repositories"
	"bones/web/actions"
	"errors"
	"log"
	"net/http"
)

type HomeContext struct {
	*BaseContext
	Users []entities.User
}

func LoadHomePage(res http.ResponseWriter, req *http.Request) {
	ctx := HomeContext{newBaseContext("index.html"), nil}

	users, err := repositories.Users().All()

	if err != nil {
		log.Println("Error loading users from repository:", err)
		ctx.AddError(errors.New("Unable to load users"))
	} else {
		ctx.Users = users
	}

	(actions.RenderPage{
		ResponseWriter: res,
		PageContext:    ctx}).Run()
}

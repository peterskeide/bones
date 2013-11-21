package handlers

import (
	"bones/entities"
	"bones/repositories"
	"bones/web/services"
	"bones/web/templating"
	"errors"
	"log"
	"net/http"
)

type HomeContext struct {
	*templating.BaseContext
	Users []entities.User
}

type HomeHandler struct {
	services.Shortcuts
	Users repositories.UserRepository
}

func (h *HomeHandler) LoadHomePage(res http.ResponseWriter, req *http.Request) {
	ctx := HomeContext{templating.NewBaseContext("index.html"), nil}
	users, err := h.Users.All()

	if err != nil {
		log.Println("Error loading users from repository:", err)
		ctx.AddError(errors.New("Unable to load users"))
	} else {
		ctx.Users = users
	}

	h.RenderPage(res, ctx)
}

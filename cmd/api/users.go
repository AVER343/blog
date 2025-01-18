package main

import (
	"net/http"

	"github.com/aver343/blog/pkg/utils"
	"github.com/go-chi/chi/v5"
)

func userHandler(app *application) func(chi.Router) {
	return func(r chi.Router) {
		r.Get("/", app.getAllUsers)
	}
}

func (app *application) getAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := app.Repository.User.GetAllUsers(ctx)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	utils.WriteJSON(w, 200, users)
}

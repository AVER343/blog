package main

import (
	"fmt"
	"net/http"

	"github.com/aver343/blog/pkg/models"
	"github.com/go-chi/chi/v5"
)

func postHandler(app *application) func(chi.Router) {
	return func(r chi.Router) {

		r.Post("/", app.createPostHandler)
		r.Get("/", app.getPostHandler)
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("CONSOLE log")
	ctx := r.Context()
	post := &models.Post{}
	app.Repository.Post.Create(ctx, post)
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {

}

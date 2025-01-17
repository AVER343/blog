package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aver343/blog/pkg/models"
	"github.com/aver343/blog/pkg/utils"
	"github.com/go-chi/chi/v5"
)

func postHandler(app *application) func(chi.Router) {
	return func(r chi.Router) {
		r.Post("/", app.createPostHandler)
		r.Get("/", app.getPostHandler)
		r.Get("/{postID}", app.getPostByIdHandler)
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("CONSOLE log")
	ctx := r.Context()
	post := &models.Post{}
	app.Repository.Post.Create(ctx, post)
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("CONSOLE log")
	ctx := r.Context()
	post := &models.Post{
		Title:   "Title",
		Content: "CONTENT",
		UserID:  "a5d5238c-8a0b-4811-9009-8fc3bfbd9154",
		Tags:    nil,
	}
	err := app.Repository.Post.Create(ctx, post)
	if err != nil {
		utils.WriteJSONError(w, 500, err.Error())
		return
	}
	utils.WriteJSON(w, 201, post)
}

func (app *application) getPostByIdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("CONSOLE log")
	postID := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ctx := context.Background()
	post, err := app.Repository.Post.GetPostByID(ctx, id)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(w, 200, post)
}

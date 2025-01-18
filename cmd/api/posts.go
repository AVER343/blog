package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/aver343/blog/pkg/models"
	"github.com/aver343/blog/pkg/utils"
	"github.com/go-chi/chi/v5"
)

func postHandler(app *application) func(chi.Router) {
	return func(r chi.Router) {
		r.Post("/", app.createPostHandler)
		r.Get("/", app.getPostsHandler)
		r.Get("/{postID}", app.getPostByIDHandler)
	}
}

func (app *application) getPostsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := app.Repository.Post.GetAllPosts(ctx)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, posts)
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	post := &models.Post{
		Title:   "Title",
		Content: "CONTENT",
		UserID:  "5ec663ac-f04e-4c8d-930b-931ed0418ef6",
		Tags:    nil,
	}
	err := app.Repository.Post.Create(ctx, post)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	utils.WriteJSON(w, 201, post)
}

func (app *application) getPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}
	ctx := context.Background()
	post, err := app.Repository.Post.GetPostByID(ctx, id)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	utils.WriteJSON(w, 200, post)
}

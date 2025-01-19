package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/aver343/blog/pkg/db/dto"
	"github.com/aver343/blog/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
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
	utils.JsonResponse(w, http.StatusOK, posts)
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Failed to read body")
		return
	}

	// Parse into struct
	var createPostParam dto.CreatePostPayload
	err = json.Unmarshal(body, &createPostParam)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	err = utils.Validate.Struct(createPostParam)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		if validationErrors != nil {
			app.badRequestError(w, r, validationErrors)
			return
		}
	}

	post, err := app.Repository.Post.Create(ctx, &createPostParam)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	utils.JsonResponse(w, 201, post)
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
	utils.JsonResponse(w, 200, post)
}

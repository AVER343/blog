package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aver343/blog/pkg/db/dto"
	"github.com/aver343/blog/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func userHandler(app *application) func(chi.Router) {
	return func(r chi.Router) {
		r.Get("/", app.getAllUsers)
		r.Get("/{userID}/posts", app.getUserPosts)
		r.Post("/register", app.register)
		r.Post("/login", app.login)
	}
}

func (app *application) getAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := app.Repository.User.GetAllUsers(ctx)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	utils.JsonResponse(w, 200, users)
}

func (app *application) getUserPosts(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	ctx := context.Background()
	postPayload := &dto.GetPostByUserIDPayload{
		UserID: userID,
	}
	post, err := app.Repository.Post.GetPostByUserID(ctx, postPayload)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	utils.JsonResponse(w, 200, post)
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Failed to read body")
		return
	}
	var userPayload dto.RegisterUserPayload
	err = json.Unmarshal(body, &userPayload)
	fmt.Print(userPayload)

	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	err = utils.Validate.Struct(userPayload)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		if validationErrors != nil {
			app.badRequestError(w, r, validationErrors)
			return
		}
		return
	}
	user, err := app.Repository.User.Create(ctx, &userPayload)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	utils.JsonResponse(w, 201, user)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Failed to read body")
		return
	}
	var userPayload dto.LoginUserPayload
	err = json.Unmarshal(body, &userPayload)
	fmt.Print(userPayload)

	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	err = utils.Validate.Struct(userPayload)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		if validationErrors != nil {
			app.badRequestError(w, r, validationErrors)
			return
		}
		return
	}
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	utils.JsonResponse(w, 201, nil)
}

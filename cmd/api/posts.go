package main

import (
	"fmt"
	"net/http"

	"github.com/PrinceNarteh/gopher-social/internal/models"
)

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload models.CreatePostDto
	if err := readJSON(w, r, &payload); err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(payload)

	post := &models.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		UserId:  1,
	}

	ctx := r.Context()
	if err := app.repo.Posts.Create(ctx, post); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

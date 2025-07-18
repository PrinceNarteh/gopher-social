package handlers

import (
	"errors"
	"net/http"

	"github.com/PrinceNarteh/social/internal/middleware"
	"github.com/PrinceNarteh/social/internal/models"
	"github.com/PrinceNarteh/social/internal/store"
	"github.com/PrinceNarteh/social/internal/utils"
)

type postKey string

const postCtx postKey = "post"

type PostHandler struct {
	store *store.Storage
}

func (h *PostHandler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var payload *models.Post
	if err := utils.ParseJSON(w, r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := utils.ValidateStruct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	payload.UserID = 1

	ctx := r.Context()
	if err := h.store.Post.Create(ctx, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteResponse(w, r, http.StatusCreated, payload)
}

func (h *PostHandler) GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := h.store.Post.GetAll(ctx)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.WriteResponse(w, r, http.StatusCreated, posts)
}

func (h *PostHandler) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	postId, err := utils.GetURLParamAsInt(r, "postId")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid post ID")
	}

	ctx := r.Context()
	post := middleware.GetPostFromCtx(r)

	comments, err := h.store.Comment.GetByPostID(ctx, postId)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	post.Comments = comments

	utils.WriteResponse(w, r, http.StatusCreated, post)
}

func (h *PostHandler) UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	postId, err := utils.GetURLParamAsInt(r, "postId")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid post ID")
		return
	}

	// var payload *models.Post
	// if err := utils.ParseJSON(w, r, &payload); err != nil {
	// 	utils.WriteError(w, http.StatusBadRequest, err.Error())
	// 	return
	// }

	utils.WriteResponse(w, r, http.StatusOK, postId)
}

func (h *PostHandler) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	postId, err := utils.GetURLParamAsInt(r, "postId")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid post ID")
		return
	}

	ctx := r.Context()
	if err := h.store.Post.Delete(ctx, postId); err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			utils.WriteError(w, http.StatusNotFound, "post not found")
		default:
			utils.InternalServerError(w, r, err)
		}
		return
	}

	utils.WriteResponse(w, r, http.StatusOK, "posts deleted successfully")
}

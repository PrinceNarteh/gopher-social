package handlers

import (
	"errors"
	"net/http"

	"github.com/PrinceNarteh/social/internal/models"
	"github.com/PrinceNarteh/social/internal/store"
	"github.com/PrinceNarteh/social/internal/utils"
)

type PostHandler struct {
	store *store.Storage
}

func getPostFromCtx(r *http.Request) *models.Post {
	if post, ok := r.Context().Value("post").(*models.Post); ok {
		return post
	}
	return nil
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

	if err := utils.WriteResponse(w, http.StatusCreated, payload); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *PostHandler) GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := h.store.Post.GetAll(ctx)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if err := utils.WriteResponse(w, http.StatusCreated, posts); err != nil {
		utils.InternalServerError(w, r, err)
	}
}

func (h *PostHandler) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	postId, err := utils.GetURLParamAsInt(r, "postId")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid post ID")
	}

	ctx := r.Context()
	post := getPostFromCtx(r)
	comments, err := h.store.Comment.GetByPostID(ctx, postId)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	post.Comments = comments

	if err := utils.WriteResponse(w, http.StatusCreated, post); err != nil {
		utils.InternalServerError(w, r, err)
	}
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

	if err := utils.WriteResponse(w, http.StatusOK, postId); err != nil {
		utils.InternalServerError(w, r, err)
	}
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

	if err := utils.WriteResponse(w, http.StatusOK, "posts deleted successfully"); err != nil {
		utils.InternalServerError(w, r, err)
	}
}

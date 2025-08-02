package handlers

import (
	"errors"
	"net/http"

	"github.com/PrinceNarteh/social/internal/middleware"
	"github.com/PrinceNarteh/social/internal/models"
	"github.com/PrinceNarteh/social/internal/store"
	"github.com/PrinceNarteh/social/internal/utils"
)

type PostHandler struct {
	store *store.Storage
}

// CreatePostHandler godoc
// @Summary      Create a new post
// @Description  Create a new post for the authenticated user
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        post  body      models.Post  true  "Post DTO"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /posts [post]
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

// GetAllPostsHandler godoc
// @Summary      Get all posts
// @Description  Retrieve all posts from the database
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        limit  query     int    false "Limit"  default(20)
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /posts [get]
func (h *PostHandler) GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := h.store.Post.GetAll(ctx)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.WriteResponse(w, r, http.StatusCreated, posts)
}

// GetPostHandler godoc
// @Summary      Get a post by ID
// @Description  Retrieve a specific post by its ID, including its comments
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        postId  path      int64  true  "Post ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /posts/{postId} [get]
func (h *PostHandler) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	post := middleware.GetPostFromCtx(r)

	comments, err := h.store.Comment.GetByPostID(ctx, post.ID)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	post.Comments = comments

	utils.WriteResponse(w, r, http.StatusOK, post)
}

// UpdatePostHandler godoc
// @Summary      Update a post by ID
// @Description  Update a specific post by its ID
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        postId  path      int64  true  "Post ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /posts/{postId} [put]
func (h *PostHandler) UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	post := middleware.GetPostFromCtx(r)

	var payload models.UpdatePostDto
	if err := utils.ParseJSON(w, r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := utils.ValidateStruct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if payload.Title != nil {
		post.Title = *payload.Title
	}

	if payload.Content != nil {
		post.Content = *payload.Content
	}

	if err := h.store.Post.Update(r.Context(), post); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.WriteResponse(w, r, http.StatusOK, post)
}

// DeletePostHandler godoc
// @Summary      Delete a post by ID
// @Description  Delete a specific post by its ID
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        postId  path      int64  true  "Post ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /posts/{postId} [delete]
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

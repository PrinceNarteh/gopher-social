package handlers

import (
	"net/http"

	"github.com/PrinceNarteh/social/internal/models"
	"github.com/PrinceNarteh/social/internal/store"
	"github.com/PrinceNarteh/social/internal/utils"
)

type FeedHandler struct {
	store *store.Storage
}

// GetUserFeedHandler godoc
// @Summary      Get user feed
// @Description  Get the feed for a user by their ID
// @Tags         feed
// @Accept       json
// @Produce      json
// @Param        id  path      int64  true  "User ID"
// @Success      200 {object}  models.PaginatedFeedResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /feed/user/{id} [get]
func (h *FeedHandler) GetUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	fq := models.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid input")
		return
	}

	if err := utils.ParseJSON(w, r, fq); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid input")
		return
	}

	posts, err := h.store.Feed.GetUserFeed(r.Context(), int64(50), fq)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.WriteResponse(w, r, http.StatusOK, posts)
}

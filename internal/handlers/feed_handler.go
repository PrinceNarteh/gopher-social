package handlers

import (
	"net/http"

	"github.com/PrinceNarteh/social/internal/store"
	"github.com/PrinceNarteh/social/internal/utils"
)

type FeedHandler struct {
	store *store.Storage
}

func (h *FeedHandler) GetUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	posts, err := h.store.Feed.GetUserFeed(ctx, int64(50))
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.WriteResponse(w, r, http.StatusOK, posts)
}

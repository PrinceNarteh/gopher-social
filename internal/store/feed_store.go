package store

import (
	"context"
	"database/sql"

	"github.com/PrinceNarteh/social/internal/models"
)

var _ FeedStore = (*feedStore)(nil)

type FeedStore interface {
	GetUserFeed(context.Context, int64) ([]models.Feed, error)
}

type feedStore struct {
	db *sql.DB
}

func (sore *feedStore) GetUserFeed(ctx context.Context, userId int64) ([]models.Feed, error) {
	var feed []models.Feed
	return feed, nil
}

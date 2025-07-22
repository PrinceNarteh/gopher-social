package store

import (
	"context"
	"database/sql"
)

var _ FollowerStore = (*followerStore)(nil)

type FollowerStore interface {
	Follow(context.Context, int64, int64) error
	Unfollow(context.Context, int64, int64) error
}

type followerStore struct {
	db *sql.DB
}

func (store *followerStore) Follow(ctx context.Context, userId, followerId int64) error {
	query := `INSERT INTO followers (user_id, follower_id) VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	if _, err := store.db.ExecContext(ctx, query, userId, followerId); err != nil {
		return err
	}

	return nil
}

func (store *followerStore) Unfollow(ctx context.Context, userId, followerId int64) error {
	query := `DELETE FROM followers WHERE user_id = $1 AND follower_id = $2`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	if _, err := store.db.ExecContext(ctx, query, userId, followerId); err != nil {
		return err
	}

	return nil
}

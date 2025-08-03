package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("record not found")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Comment  CommentStore
	Feed     FeedStore
	Follower FollowerStore
	Post     PostStore
	User     UserStore
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Comment:  &commentStore{db: db},
		Feed:     &feedStore{db: db},
		Follower: &followerStore{db: db},
		Post:     &postStore{db: db},
		User:     &userStore{db: db},
	}
}

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err = fn(tx); err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}

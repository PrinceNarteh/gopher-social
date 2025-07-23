package store

import (
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

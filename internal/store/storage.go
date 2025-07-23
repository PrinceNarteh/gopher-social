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
	Post     PostStore
	User     UserStore
	Comment  CommentStore
	Follower FollowerStore
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Post:     &postStore{db: db},
		User:     &userStore{db: db},
		Comment:  &commentStore{db: db},
		Follower: &followerStore{db: db},
	}
}

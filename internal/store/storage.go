package store

import (
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("record not found")

type Storage struct {
	Post    PostStore
	User    UserStore
	Comment CommentStore
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Post:    &postStore{db: db},
		User:    &userStore{db: db},
		Comment: &commentStore{db: db},
	}
}

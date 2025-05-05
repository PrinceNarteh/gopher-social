package repositories

import (
	"context"
	"database/sql"

	"github.com/PrinceNarteh/gopher-social/internal/models"
)

type Repositories struct {
	Posts interface {
		Create(context.Context, *models.Post) error
	}
	Users interface {
		Create(context.Context, *models.User) error
	}
}

func NewRepositories(db *sql.DB) Repositories {
	return Repositories{
		Posts: &PostRespository{db: db},
		Users: &UserRepository{db: db},
	}
}

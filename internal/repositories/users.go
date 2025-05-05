package repositories

import (
	"context"
	"database/sql"

	"github.com/PrinceNarteh/gopher-social/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (first_name, last_name, username, email, password)
		VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at
	`

	if err := r.db.QueryRowContext(
		ctx,
		query,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
	).Scan(
		&user.Id,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return err
	}

	return nil
}

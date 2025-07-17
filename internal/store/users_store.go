package store

import (
	"context"
	"database/sql"

	"github.com/PrinceNarteh/social/internal/models"
)

type UserStore interface {
	Create(context.Context, *models.User) error
	FindByEmail(context.Context, string) (*models.User, error)
}

type UserStoreImpl struct {
	db *sql.DB
}

func (store *UserStoreImpl) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO 
		users (first_name, last_name, username, email, password)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	if err := store.db.QueryRowContext(
		ctx,
		query,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return err
	}

	return nil
}

func (store *UserStoreImpl) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT * FROM users WHERE email = $1
	`

	user := new(models.User)
	if err := store.db.QueryRowContext(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return user, nil
}

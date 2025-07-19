package store

import (
	"context"
	"database/sql"

	"github.com/PrinceNarteh/social/internal/models"
)

var _ UserStore = (*userStore)(nil)

type UserStore interface {
	FindById(context.Context, int64) (*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	Create(context.Context, *models.User) error
	Update(context.Context, *models.UpdateUserDto) error
	Delete(context.Context, int64) error
}

type userStore struct {
	db *sql.DB
}

func (store *userStore) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO 
		users (first_name, last_name, username, email, password)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

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

func (store *userStore) findUser(ctx context.Context, query string, args ...any) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := new(models.User)
	if err := store.db.QueryRowContext(
		ctx,
		query,
		args...,
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
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}

func (store *userStore) FindById(ctx context.Context, id int64) (*models.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, password, created_at, updated_at 
		FROM users 
		WHERE id = $1
	`
	return store.findUser(ctx, query, id)
}

func (store *userStore) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, password, created_at, updated_at 
		FROM users 
		WHERE email = $1
	`
	return store.findUser(ctx, query, email)
}

func (store *userStore) Update(ctx context.Context, payload *models.UpdateUserDto) error {
	query := `
	UPADATE users
	SET first_name, last_name, username
	VALUES $1, $2, $3
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	store.db.ExecContext(ctx, query)

	return nil
}

func (store *userStore) Delete(context.Context, int64) error {
	return nil
}

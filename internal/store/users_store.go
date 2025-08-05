package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/PrinceNarteh/social/internal/models"
	"github.com/PrinceNarteh/social/internal/utils"
)

var (
	ErrDuplicateEmail    = errors.New("a user with that email already exists")
	ErrDuplicateUsername = errors.New("a user with that username already exists")
)

var _ UserStore = (*userStore)(nil)

type UserStore interface {
	Activate(context.Context, string) error
	Create(context.Context, *sql.Tx, *models.User) error
	CreateAndInvite(context.Context, *models.User, string, time.Duration) error
	FindById(context.Context, int64) (*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	Delete(context.Context, int64) error
}

type userStore struct {
	db *sql.DB
}

func (s *userStore) Activate(ctx context.Context, token string) error {
	return withTx(s.db, ctx, func(tx *sql.Tx) error {
		// 1. find the user that this token belongs to
		user, err := s.getUserFromInvitation(ctx, tx, token)
		if err != nil {
			return err
		}

		// 2. update the user is_active field to true
		user.IsActive = true
		if err := s.updateUser(ctx, tx, user); err != nil {
			return err
		}

		// 3. Delete user invitations
		if err := s.deleteUserInvitations(ctx, tx, user.ID); err != nil {
			return err
		}

		return nil
	})
}

func (s *userStore) Create(ctx context.Context, tx *sql.Tx, user *models.User) error {
	query := `
		INSERT INTO 
		users (first_name, last_name, username, email, password)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	if err := tx.QueryRowContext(
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
		switch err.Error() {
		case `pq: duplicate key value violates unique constraint "users_email_key"`:
		case `pq: duplicate key value violates unique constraint "users_username_key"`:
		default:
			return err
		}
	}

	return nil
}

func (s *userStore) CreateAndInvite(ctx context.Context, user *models.User, token string, exp time.Duration) error {
	return withTx(s.db, ctx, func(tx *sql.Tx) error {
		if err := s.Create(ctx, tx, user); err != nil {
			return err
		}

		if err := s.createUserInvitation(ctx, tx, user.ID, token, exp); err != nil {
			return err
		}

		return nil
	})
}

func (s *userStore) Delete(ctx context.Context, userId int64) error {
	return withTx(s.db, ctx, func(tx *sql.Tx) error {
		if err := s.deleteUser(ctx, tx, userId); err != nil {
			return nil
		}

		if err := s.deleteUserInvitations(ctx, tx, userId); err != nil {
			return nil
		}

		return nil
	})
}

func (s *userStore) findUser(ctx context.Context, query string, args ...any) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := new(models.User)
	if err := s.db.QueryRowContext(
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

func (s *userStore) FindById(ctx context.Context, id int64) (*models.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, password, created_at, updated_at 
		FROM users 
		WHERE id = $1
	`
	return s.findUser(ctx, query, id)
}

func (s *userStore) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, password, created_at, updated_at 
		FROM users 
		WHERE email = $1
	`
	return s.findUser(ctx, query, email)
}

func (s *userStore) createUserInvitation(
	ctx context.Context,
	tx *sql.Tx,
	userId int64,
	token string,
	exp time.Duration,
) error {
	query := `INSERT INTO user_invitations (user_id, token, expiry) VALUES ($1, $2, $3)`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	if _, err := tx.ExecContext(ctx, query, userId, token, time.Now().Add(exp)); err != nil {
		return err
	}

	return nil
}

func (s *userStore) updateUser(ctx context.Context, tx *sql.Tx, user *models.User) error {
	query := `
		UPADATE users
		SET is_active = $1 
		WHERE id = $2
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	if _, err := tx.ExecContext(ctx, query, user.IsActive, user.ID); err != nil {
		return err
	}

	return nil
}

func (s *userStore) deleteUser(ctx context.Context, tx *sql.Tx, userId int64) error {
	query := `DELETE FROM users WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	if _, err := tx.ExecContext(ctx, query, userId); err != nil {
		return err
	}

	return nil
}

func (s *userStore) getUserFromInvitation(ctx context.Context, tx *sql.Tx, token string) (*models.User, error) {
	query := `
	 	SELECT u.id, u.first_name, u.last_name, u.username, u.email, u.is_active, u.created_at, u.updated_at
		FROM users u 
		JOIN user_invitations ui ON u.id = ui.user_id
		WHERE u.token = $1 and ui.expiry > $2
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &models.User{}
	hashedToken := utils.HashToken(token)
	if err := tx.QueryRowContext(ctx, query, hashedToken, time.Now()).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.IsActive,
		&user.CreatedAt,
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

func (s *userStore) deleteUserInvitations(ctx context.Context, tx *sql.Tx, userId int64) error {
	query := `DELETE FROM user_invitations WHERE user_id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	if _, err := tx.ExecContext(ctx, query, userId); err != nil {
		return err
	}

	return nil
}

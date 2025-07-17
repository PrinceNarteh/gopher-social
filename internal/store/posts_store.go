package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"

	"github.com/PrinceNarteh/social/internal/models"
)

type PostStore interface {
	Create(context.Context, *models.Post) error
	Delete(context.Context, int64) error
	GetByID(ctx context.Context, id int64) (*models.Post, error)
	GetAll(ctx context.Context) ([]models.Post, error)
}

type PostStoreImpl struct {
	db *sql.DB
}

func (s *PostStoreImpl) Create(ctx context.Context, post *models.Post) error {
	query := `
		INSERT INTO posts 
		(title, content, tags, user_id)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	if err := s.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		pq.Array(post.Tags),
		post.UserID,
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	); err != nil {
		return err
	}

	return nil
}

func (s *PostStoreImpl) GetAll(ctx context.Context) ([]models.Post, error) {
	query := `SELECT id, title, content, tags, user_id, created_at, updated_at FROM posts`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			pq.Array(&post.Tags),
			&post.UserID,
			&post.CreatedAt,
			&post.UpdatedAt,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (s *PostStoreImpl) GetByID(ctx context.Context, id int64) (*models.Post, error) {
	query := `
		SELECT id, title, content, tags, user_id, created_at, updated_at
		FROM posts 
		WHERE id = $1
	`

	post := new(models.Post)
	if err := s.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		pq.Array(&post.Tags),
		&post.UserID,
		&post.CreatedAt,
		&post.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return post, nil
}

func (s *PostStoreImpl) Delete(ctx context.Context, postId int64) error {
	query := `DELETE FROM posts WHERE id = $1`

	res, err := s.db.ExecContext(ctx, query, postId)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

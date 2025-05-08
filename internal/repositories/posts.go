package repositories

import (
	"context"
	"database/sql"

	"github.com/lib/pq"

	"github.com/PrinceNarteh/gopher-social/internal/models"
)

type PostRespository struct {
	db *sql.DB
}

func (r *PostRespository) Create(ctx context.Context, post *models.Post) error {
	query := `
		INSERT INTO post (title, content, tags, user_id)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	if err := r.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		pq.Array(post.Tags),
		post.UserId,
	).Scan(
		&post.Id,
		&post.CreatedAt,
		&post.UpdatedAt,
	); err != nil {
		return err
	}

	return nil
}

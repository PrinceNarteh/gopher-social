package store

import (
	"context"
	"database/sql"

	"github.com/PrinceNarteh/social/internal/models"
)

var _ CommentStore = (*commentStore)(nil)

type CommentStore interface {
	GetByPostID(context.Context, int64) (*[]models.Comment, error)
}

type commentStore struct {
	db *sql.DB
}

func (store *commentStore) GetByPostID(ctx context.Context, postID int64) (*[]models.Comment, error) {
	query := `
		SELECT c.id, c.post_id, c.user_id, c.content, c.created_at, u.id, u.username  
		FROM comments AS c
		JOIN users AS u ON c.user_id = u.id 
		WHERE c.post_id = $1
		ORDER BY c.created_at DESC
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := store.db.QueryContext(ctx, query, postID)
	if err == nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var c models.Comment
		c.User = models.User{}
		if err := rows.Scan(
			&c.ID,
			&c.PostID,
			&c.UserID,
			&c.Content,
			&c.CreatedAt,
			&c.User.ID,
			&c.User.Username,
		); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return &comments, nil
}

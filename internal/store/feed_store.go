package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"

	"github.com/PrinceNarteh/social/internal/models"
)

var _ FeedStore = (*feedStore)(nil)

type FeedStore interface {
	GetUserFeed(context.Context, int64, models.PaginatedFeedQuery) ([]models.Feed, error)
}

type feedStore struct {
	db *sql.DB
}

func (store *feedStore) GetUserFeed(
	ctx context.Context,
	userId int64,
	fq models.PaginatedFeedQuery,
) ([]models.Feed, error) {
	query := `
		SELECT
			p.id, p.user_id, p.title, p.content, p.created_at, p.version, p.tags,
			u.username,
			COUNT(c.id) AS comments_count
		FROM posts AS p
		LEFT JOIN comments AS c ON c.post_id = p.id
		LEFT JOIN users AS u ON p.user_id = u.id
		JOIN followers AS f ON f.follwer_id = p.user_id OR p.user_id = $1
		WHERE f.user_id = 341 OR p.user_id = $1
		GROUP BY p.id, u.username
		ORDER BY p.created_at $2
		LIMIT $3 OFFSET $4
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := store.db.QueryContext(ctx, query, userId, fq.Sort, fq.Limit, fq.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feed []models.Feed
	for rows.Next() {
		var f models.Feed
		if err != rows.Scan(
			&f.ID,
			&f.UserID,
			&f.Title,
			&f.Content,
			&f.CreatedAt,
			&f.Versoin,
			pq.Array(&f.Tags),
			&f.User.Username,
			&f.CommentCount,
		) {
			return nil, err
		}

		feed = append(feed, f)
	}

	return feed, nil
}

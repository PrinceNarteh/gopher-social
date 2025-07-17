package models

type Comment struct {
	ID        int64  `json:"id,omitempty"`
	PostID    int64  `json:"post_id"              validate:"required,gte=1"`
	UserID    int64  `json:"user_id"              validate:"required,gte=2"`
	Content   string `json:"content"              validate:"required"`
	CreatedAt string `json:"created_at,omitempty"`
	User      User   `json:"user"`
}

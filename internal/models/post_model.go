package models

type Post struct {
	ID        int64      `json:"id,omitempty"`
	Title     string     `json:"title"               validate:"required"`
	Content   string     `json:"content"             validate:"required"`
	UserID    int64      `json:"userId"              validate:"required,gte=1"`
	Tags      []string   `json:"tags"                validate:"required"`
	Comments  *[]Comment `json:"comments"`
	CreatedAt string     `json:"createdAt,omitempty"`
	UpdatedAt string     `json:"updatedAt,omitempty"`
}

type UpdatePostDto struct {
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content"`
}

package models

type Feed struct {
	Post
	CommentCount int `json:"commentCount"`
}

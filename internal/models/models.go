// Package models provides the models for the application
package models

type Models struct {
	Feed *Feed
	Post *Post
	User *User
}

func NewModels() *Models {
	return &Models{
		Feed: &Feed{},
		Post: &Post{},
		User: &User{},
	}
}

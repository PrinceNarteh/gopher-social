// Package models provides the models for the application
package models

type Models struct {
	Post *Post
	User *User
}

func NewModels() *Models {
	return &Models{
		Post: &Post{},
		User: &User{},
	}
}

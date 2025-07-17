package models

type User struct {
	ID        int64  `json:"id,omitempty"`
	FirstName string `json:"firstName"    validate:"required"`
	LastName  string `json:"lastName"     validate:"required"`
	Username  string `json:"username"     validate:"required"`
	Email     string `json:"email"        validate:"required,email"`
	Password  string `json:"password"     validate:"required,min=6"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

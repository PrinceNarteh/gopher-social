package models

type User struct {
	ID        int64  `json:"id,omitempty"`
	FirstName string `json:"firstName"    validate:"required"`
	LastName  string `json:"lastName"     validate:"required"`
	Username  string `json:"username"     validate:"required"`
	Email     string `json:"email"        validate:"required,email"`
	Password  string `json:"-"            validate:"required,min=6"`
	IsActive  bool   `json:"isActive"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserWithToken struct {
	User  User
	Token string
}

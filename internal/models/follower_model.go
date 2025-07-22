package models

type Follower struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"userId"`
	FollowerID int64  `json:"follower_id"`
	CreatedAt  string `json:"created_at"`
}

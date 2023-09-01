package models

import "time"

type User struct {
	UserId         int       `db:"user_id"`
	Username       string    `db:"username"`
	HashedPassword string    `db:"hashed_password"`
	Role           string    `db:"role"`
	CreatedAt      time.Time `db:"created_at"`
	LastLogin      time.Time `db:"last_login"`
}

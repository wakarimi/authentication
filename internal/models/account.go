package models

import "time"

type Account struct {
	AccountId      int       `db:"account_id"`
	Username       string    `db:"username"`
	HashedPassword string    `db:"hashed_password"`
	CreatedAt      time.Time `db:"created_at"`
	LastLogin      time.Time `db:"last_login"`
}

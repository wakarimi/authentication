package account

import "time"

type Account struct {
	ID             int        `db:"id"`
	Username       string     `db:"username"`
	HashedPassword string     `db:"hashed_password"`
	CreatedAt      time.Time  `db:"created_at"`
	LastSignIn     *time.Time `db:"last_sign_in"`
}

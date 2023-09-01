package models

import "time"

type Token struct {
	TokenId  int       `db:"token_id"`
	UserId   int       `db:"user_id"`
	Type     string    `db:"type"`
	Value    string    `db:"value"`
	ExpiryAt time.Time `db:"expiry_at"`
}

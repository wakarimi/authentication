package model

import "time"

type RefreshToken struct {
	ID        int       `db:"id"`
	DeviceID  int       `db:"device_id"`
	Token     string    `db:"token"`
	CreatedAt time.Time `db:"created_at"`
	ExpiresAt time.Time `db:"expires_at"`
}

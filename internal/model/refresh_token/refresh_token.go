package refresh_token

import "time"

const (
	Duration = time.Hour * 24 * 14
)

type Payload struct {
	AccountID int   `json:"accountId"`
	DeviceID  int   `json:"deviceId"`
	IssuedAt  int64 `json:"issuedAt"`
	ExpiryAt  int64 `json:"expiryAt"`
}

type RefreshToken struct {
	ID        int       `db:"id"`
	DeviceID  int       `db:"device_id"`
	Token     string    `db:"token"`
	CreatedAt time.Time `db:"created_at"`
	ExpiresAt time.Time `db:"expires_at"`
}

package access_token

import "time"

const (
	Duration = time.Minute * 10
)

type Payload struct {
	AccountID int      `json:"accountId"`
	DeviceID  int      `json:"deviceId"`
	Roles     []string `json:"roles"`
	IssuedAt  int64    `json:"issuedAt"`
	ExpiryAt  int64    `json:"expiryAt"`
}

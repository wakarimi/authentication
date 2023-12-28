package token

type RefreshTokenPayload struct {
	AccountID int   `json:"accountId"`
	DeviceID  int   `json:"deviceId"`
	IssuedAt  int64 `json:"issuedAt"`
	ExpiryAt  int64 `json:"expiryAt"`
}

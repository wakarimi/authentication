package types

// TokenPayload represents the decoded claims from a JWT token.
type TokenPayload struct {
	// AccountID represents the ID of the associated account.
	// Required: true
	// Example: 12345
	AccountID int `json:"account_id"`

	// ExpiryAt represents the UNIX timestamp when the token will expire.
	// Required: true
	// Example: 1615569457
	ExpiryAt int64 `json:"expiry_at"`

	// Type indicates the type of the token (e.g., "access" or "refresh").
	// Required: true
	// Example: "access"
	Type string `json:"type"`
}

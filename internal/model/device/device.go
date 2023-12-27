package device

type Device struct {
	ID          int    `db:"id"`
	AccountID   int    `db:"account_id"`
	Fingerprint string `db:"fingerprint"`
}

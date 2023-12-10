package model

type Device struct {
	ID          int    `db:"id"`
	AccountID   int    `db:"account_id"`
	Name        string `db:"name"`
	Fingerprint string `db:"fingerprint"`
}

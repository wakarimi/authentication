package models

type AccountRole struct {
	AccountId int    `db:"account_id"`
	Role      string `db:"role"`
}

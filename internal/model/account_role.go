package model

type AccountRole struct {
	AccountID int      `db:"account_id"`
	Role      RoleName `db:"role"`
}

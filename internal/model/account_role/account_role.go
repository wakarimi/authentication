package account_role

type RoleName string

const (
	RoleAdmin RoleName = "ADMIN"
	RoleUser  RoleName = "USER"
)

type AccountRole struct {
	AccountID int      `db:"account_id"`
	Role      RoleName `db:"role"`
}

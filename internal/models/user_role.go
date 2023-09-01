package models

type UserRole struct {
	UserId int `db:"user_id"`
	RoleId int `db:"role_id"`
}

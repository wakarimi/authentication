package models

type Role struct {
	RoleId int    `db:"role_id"`
	Name   string `db:"name"`
}

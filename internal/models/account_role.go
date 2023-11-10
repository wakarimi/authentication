package models

type AccountRole struct {
	Id   int    `db:"id"`
	Role string `db:"role"`
}

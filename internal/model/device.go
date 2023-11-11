package model

type Device struct {
	Id        int    `db:"id"`
	AccountId int    `db:"account_id"`
	Name      string `db:"name"`
}

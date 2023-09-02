package repository

import (
	"authentication/internal/database"
	"authentication/internal/models"
	"github.com/jmoiron/sqlx"
)

func AccountExist(username string) (bool, error) {
	var count int

	query := `
		SELECT COUNT(account_id)
		FROM accounts
		WHERE username = $1
	`
	err := database.Db.Get(
		&count, query, username)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func FetchAccountByUsername(username string) (models.Account, error) {
	var account models.Account

	query := `
		SELECT account_id, username, hashed_password, created_at, last_login
		FROM accounts
		WHERE username = $1
	`
	err := database.Db.Get(&account, query, username)
	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}

func CreateAccountTx(tx *sqlx.Tx, account models.Account) (int, error) {
	var accountId int

	query := `
		INSERT INTO accounts(username, hashed_password) 
		VALUES ($1, $2)
		RETURNING account_id
	`
	err := tx.QueryRow(query, account.Username, account.HashedPassword).Scan(&accountId)
	if err != nil {
		return 0, err
	}

	return accountId, nil
}

func CountAccount() (int, error) {
	var count int

	query := `
		SELECT COUNT(*)
		FROM accounts
	`
	err := database.Db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

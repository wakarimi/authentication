package repository

import (
	"authentication/internal/database"
	"authentication/internal/models"
)

func UserExist(username string) (bool, error) {
	var count int
	err := database.Db.Get(
		&count, `
			SELECT COUNT(user_id)
			FROM users
			WHERE username = $1
		`, username)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateUser(user models.User) error {
	query := `
		INSERT INTO users (username, hashed_password, role)
		VALUES (:username, :hashed_password, :role)
	`
	_, err := database.Db.NamedExec(query, user)
	return err
}

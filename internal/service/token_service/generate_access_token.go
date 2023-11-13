package token_service

import "github.com/jmoiron/sqlx"

func (s Service) GenerateAccessToken(tx *sqlx.Tx, refreshToken string) (accessToken string, err error) {
	return "IMPLEMENT ME", nil
}

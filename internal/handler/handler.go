package handler

import "github.com/nicksnyder/go-i18n/v2/i18n"

type useCase interface {
	SignUp(username string, password string) error
	SignIn(username string, password string, fingerprint string) (refreshToken string, accessToken string, err error)
	RefreshTokens(oldRefreshToken string) (refreshToken string, accessToken string, err error)
	AssignRole(requesterID int, accountID int, roleName string) error
}

type Handler struct {
	useCase      useCase
	bundle       i18n.Bundle
	engLocalizer i18n.Localizer
}

func New(useCase useCase,
	bundle i18n.Bundle) *Handler {
	return &Handler{
		useCase:      useCase,
		bundle:       bundle,
		engLocalizer: *i18n.NewLocalizer(&bundle, "en_US"),
	}
}

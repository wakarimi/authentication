package handler

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"wakarimi-authentication/internal/model/access_token"
	"wakarimi-authentication/internal/model/account"
	"wakarimi-authentication/internal/model/account_role"
)

type useCase interface {
	SignUp(username string, password string) error
	SignIn(username string, password string, fingerprint string) (refreshToken string, accessToken string, err error)
	RefreshTokens(oldRefreshToken string) (refreshToken string, accessToken string, err error)
	AssignRole(requesterID int, accountID int, roleName string) error
	ChangePassword(accountID int, oldPassword string, newPassword string) error
	SignOut(deviceID int) error
	SignOutAll(accountID int) error
	RevokeRole(requesterID int, accountID int, roleName string) error
	GetAccountDetails(requesterID int, accountID int) (acc account.Account, roles []account_role.AccountRole, err error)
	GetAccounts(ids []int) ([]account.Account, error)
	VerifyAccessToken(accessToken string) (bool, access_token.Payload)
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

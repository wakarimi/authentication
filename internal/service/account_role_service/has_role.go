package account_role_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account_role"
)

func (s Service) HasRole(tx *sqlx.Tx, accountID int, roleName account_role.RoleName) (bool, error) {
	roles, err := s.accountRoleRepo.ReadAllByAccount(tx, accountID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read account's roles")
		return false, err
	}

	for _, role := range roles {
		if role.Role == roleName {
			return true, nil
		}
	}

	return false, nil
}

package account_role_service

import (
	"fmt"
	"wakarimi-authentication/internal/model/account_role"
)

func (s Service) StringToRole(roleAsString string) (account_role.RoleName, error) {
	switch roleAsString {
	case "ADMIN":
		return account_role.RoleAdmin, nil
	case "USER":
		return account_role.RoleUser, nil
	default:
		err := fmt.Errorf("role %s not found", roleAsString)
		return account_role.RoleUser, err
	}
}

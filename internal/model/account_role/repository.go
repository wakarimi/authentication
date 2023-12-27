package account_role

import (
	"wakarimi-authentication/internal/service"
)

type Repository interface {
	service.Transactor
}

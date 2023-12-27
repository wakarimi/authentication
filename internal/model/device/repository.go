package device

import (
	"wakarimi-authentication/internal/service"
)

type Repository interface {
	service.Transactor
}

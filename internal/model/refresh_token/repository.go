package refresh_token

import (
	"wakarimi-authentication/internal/service"
)

type Repository interface {
	service.Transactor
}

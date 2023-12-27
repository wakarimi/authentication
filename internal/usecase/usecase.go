package usecase

type accountRoleService interface {
}

type accountService interface {
}

type deviceService interface {
}

type refreshTokenService interface {
}

type UseCase struct {
	accountRoleService  accountRoleService
	accountService      accountService
	deviceService       deviceService
	refreshTokenService refreshTokenService
}

func NewUseCase(accountRoleService accountRoleService,
	accountService accountService,
	deviceService deviceService,
	refreshTokenService refreshTokenService) *UseCase {
	return &UseCase{
		accountRoleService:  accountRoleService,
		accountService:      accountService,
		deviceService:       deviceService,
		refreshTokenService: refreshTokenService,
	}
}

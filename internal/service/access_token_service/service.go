package access_token_service

type Service struct {
	secretKey string
}

func New(secretKey string) *Service {
	return &Service{
		secretKey: secretKey,
	}
}

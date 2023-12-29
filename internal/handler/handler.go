package handler

type useCase interface {
}

type Handler struct {
	useCase useCase
}

func New(useCase useCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

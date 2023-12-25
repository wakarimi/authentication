package v1

import "github.com/rs/zerolog"

type UseCase interface {
}

type Handler struct {
	uc     UseCase
	logger zerolog.Logger
}

func NewHandler(uc UseCase, logger zerolog.Logger) *Handler {
	return &Handler{
		uc:     uc,
		logger: logger,
	}
}

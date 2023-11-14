package account_service

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) HashPassword(password string) (hashedPassword string, err error) {
	log.Debug().Msg("Password hashing")

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("Failed to hash password")
		return "", err
	}
	hashedPassword = string(hashedBytes)

	log.Debug().Msg("Password hashed")
	return hashedPassword, nil
}

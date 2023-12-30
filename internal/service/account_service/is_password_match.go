package account_service

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) IsPasswordMatch(password string, hashedPassword string) bool {
	log.Debug().Msg("Checking password matching")
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

package utils

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (hashedPassword string, err error) {
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

func CheckPasswordHash(password string, hash string) (isMatch bool) {
	log.Debug().Msg("Checking password hash")
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	log.Debug().Bool("isMatch", err == nil).Msg("Password hash checked")
	return err == nil
}

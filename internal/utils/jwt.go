package utils

import (
	"authentication/internal/config"
	"errors"
	"github.com/form3tech-oss/jwt-go"
	"time"
)

const (
	AccessTokenDuration   = time.Minute * 10
	RefreshTokenDuration  = time.Hour * 24 * 7
	RefreshTokenThreshold = time.Hour * 24 * 3
	AccessTokenType       = "access"
	RefreshTokenType      = "refresh"
)

func GenerateTokens(cfg *config.Configuration, accountId int) (refreshToken string, accessToken string, err error) {
	refreshToken, err = generateToken(cfg.RefreshSecretKey, accountId, RefreshTokenDuration, RefreshTokenType)
	if err != nil {
		return "", "", err
	}

	accessToken, err = generateToken(cfg.AccessSecretKey, accountId, AccessTokenDuration, AccessTokenType)
	if err != nil {
		return "", "", err
	}

	return refreshToken, accessToken, nil
}

func RefreshTokens(cfg *config.Configuration, refreshToken string) (newRefreshToken string, accessToken string, err error) {
	token, err := validateToken(cfg, refreshToken, RefreshTokenType)
	if err != nil {
		return "", "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	accountId := int(claims["account_id"].(float64))

	accessToken, err = generateToken(cfg.AccessSecretKey, accountId, AccessTokenDuration, AccessTokenType)
	if err != nil {
		return "", "", err
	}

	if time.Unix(int64(claims["expiry_at"].(float64)), 0).Sub(time.Now()) < RefreshTokenThreshold {
		newRefreshToken, err = generateToken(cfg.RefreshSecretKey, accountId, RefreshTokenDuration, RefreshTokenType)
		if err != nil {
			return "", "", err
		}
		return accessToken, newRefreshToken, nil
	}

	return "", accessToken, nil
}

func ValidateToken(cfg *config.Configuration, tokenString string, tokenType string) (*jwt.Token, error) {
	return validateToken(cfg, tokenString, tokenType)
}

func generateToken(secretKey string, accountId int, duration time.Duration, tokenType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account_id": accountId,
		"expiry_at":  time.Now().Add(duration).Unix(),
		"type":       tokenType,
	})

	return token.SignedString([]byte(secretKey))
}

func validateToken(cfg *config.Configuration, tokenString string, tokenType string) (*jwt.Token, error) {
	var secretKey string
	if tokenType == AccessTokenType {
		secretKey = cfg.AccessSecretKey
	} else if tokenType == RefreshTokenType {
		secretKey = cfg.RefreshSecretKey
	} else {
		return nil, errors.New("invalid token type")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if int64(time.Now().Unix()) > int64(claims["expiry_at"].(float64)) {
			return nil, errors.New("token has expired")
		}
		return token, nil
	}

	return nil, errors.New("invalid token")
}

package handlers

import (
	"authentication/internal/config"
	"authentication/internal/handlers/types"
	"authentication/internal/utils"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ValidateTokenRequest represents the payload required to validate an access token.
type ValidateTokenRequest struct {
	// Token represents the access token to be validated.
	// Required: true
	// Example: aaaaa.bbbbb.ccccc
	Token string `json:"token" binding:"required"`
}

// ValidateTokenResponse represents the response after token validation.
type ValidateTokenResponse struct {
	// Valid indicates whether the token is valid or not.
	// Required: true
	Valid bool `json:"valid"`

	// Payload contains the decoded claims from the token.
	Payload *types.TokenPayload `json:"payload,omitempty"`

	// Error contains any potential error messages during validation.
	// Example: Invalid request body.
	Error string `json:"error,omitempty"`
}

// ValidateToken
// @Summary Validate access token
// @Description Check if the provided access token is valid.
// @Accept  json
// @Produce  json
// @Param input body ValidateTokenRequest true "Token payload"
// @Success 200 {object} ValidateTokenResponse "Ok."
// @Failure 400,401,500 {object} types.Error "Error."
// @Router /validate-token [post]
func ValidateToken(c *gin.Context, cfg *config.Configuration) {
	var request ValidateTokenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ValidateTokenResponse{
			Valid: false,
			Error: "Invalid request body.",
		})
		return
	}

	var token *jwt.Token
	token, err := utils.ValidateToken(cfg, request.Token, utils.AccessTokenType)
	if err != nil {
		token, err = utils.ValidateToken(cfg, request.Token, utils.RefreshTokenType)
		if err != nil {
			c.JSON(http.StatusUnauthorized, ValidateTokenResponse{
				Valid: false,
				Error: "Invalid or expired token.",
			})
			return
		}
	}

	claimsMap := token.Claims.(jwt.MapClaims)
	tokenPayload, err := utils.MapClaimsToTokenPayload(claimsMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ValidateTokenResponse{
			Valid: false,
			Error: "Failed to process token payload.",
		})
		return
	}

	c.JSON(http.StatusOK, ValidateTokenResponse{
		Valid:   true,
		Payload: tokenPayload,
	})
}

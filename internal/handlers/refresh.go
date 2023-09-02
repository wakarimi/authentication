package handlers

import (
	"authentication/internal/config"
	"authentication/internal/handlers/types"
	"authentication/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RefreshRequest represents the payload required to refresh access token.
type RefreshRequest struct {
	// Token used to refresh the access token.
	// Required: true
	// Example: aaaaa.bbbbb.ccccc
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// RefreshResponse represents the tokens received after successful token refresh.
type RefreshResponse struct {
	// RefreshToken used to renew the access token. This may be omitted if not needed.
	// Example: aaaaa.bbbbb.ccccc
	RefreshToken string `json:"refreshToken,omitempty"`

	// AccessToken used to authenticate subsequent requests.
	// Required: true
	// Example: aaaaa.bbbbb.ccccc
	AccessToken string `json:"accessToken" binding:"required"`
}

// Refresh
// @Summary Refresh access token
// @Description Refresh access token using refresh token.
// @Accept  json
// @Produce  json
// @Param input body RefreshRequest true "Refresh token payload"
// @Success 200 {object} RefreshResponse "Ok."
// @Failure 400,401,500 {object} types.Error "Error."
// @Router /refresh [post]
func Refresh(c *gin.Context, cfg *config.Configuration) {
	var request RefreshRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, types.Error{
			Error: err.Error(),
		})
		return
	}

	_, err := utils.ValidateToken(cfg, request.RefreshToken, utils.RefreshTokenType)
	if err != nil {
		c.JSON(http.StatusUnauthorized, types.Error{
			Error: "Invalid refresh token",
		})
		return
	}

	refreshToken, accessToken, err := utils.RefreshTokens(cfg, request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.Error{
			Error: "Failed to generate tokens",
		})
		return
	}

	c.JSON(http.StatusOK, RefreshResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	})
}

package handlers

import (
	"authentication/internal/config"
	"authentication/internal/database/repository"
	"authentication/internal/handlers/types"
	"authentication/internal/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginRequest represents the payload required to login into account.
type LoginRequest struct {
	// Username of the account to be login.
	// Required: true
	// Example: Zalimannard
	Username string `json:"username" binding:"required,alphanum,min=3,max=30"`

	// Password of the account to be login.
	// Required: true
	// Example: qwerty01
	Password string `json:"password" binding:"required,alphanum,min=8,max=50"`
}

// LoginResponse represents the tokens received after successful authentication.
type LoginResponse struct {
	// RefreshToken used to renew the access token.
	// Required: true
	// Example: aaaaa.bbbbb.ccccc
	RefreshToken string `json:"refreshToken" binding:"required"`

	// AccessToken used to authenticate subsequent requests.
	// Required: true
	// Example: aaaaa.bbbbb.ccccc
	AccessToken string `json:"accessToken" binding:"required"`
}

// Login
// @Summary Login into account
// @Description Login into account with the input payload.
// @Accept  json
// @Produce  json
// @Param input body LoginRequest true "Login payload"
// @Success 200 {object} LoginResponse "Ok."
// @Failure 400,401,500 {object} types.Error "Error."
// @Router /login [post]
func Login(c *gin.Context, cfg *config.Configuration) {
	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, types.Error{
			Error: err.Error(),
		})
		return
	}

	account, err := repository.FetchAccountByUsername(request.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, types.Error{
				Error: "Invalid login credentials",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, types.Error{
			Error: "Internal server error",
		})
		return
	}

	equal := utils.CheckPasswordHash(request.Password, account.HashedPassword)
	if !equal {
		c.JSON(http.StatusUnauthorized, types.Error{
			Error: "Invalid login credentials",
		})
		return
	}

	refreshToken, accessToken, err := utils.GenerateTokens(cfg, account.AccountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.Error{
			Error: "Failed to generate tokens",
		})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	})
}

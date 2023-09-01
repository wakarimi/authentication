package handlers

import (
	"authentication/internal/config"
	"authentication/internal/database/repository"
	"authentication/internal/handlers/types"
	"authentication/internal/models"
	"authentication/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterMicroserviceRequest represents the payload required to register a new user.
type RegisterMicroserviceRequest struct {
	// Username of the user to be registered.
	// Required: true
	// Example: Music
	Username string `json:"username" binding:"required,alphanum,min=3,max=30"`

	// Password of the user to be registered.
	// Required: true
	// Example: astoyuzcnvoiaersparsarstaoeizxcnvarst
	Password string `json:"password" binding:"required,alphanum,min=8,max=50"`
}

// @Summary Register a new microservice
// @Description Register a new microservice with the input payload
// @Accept  json
// @Produce  json
// @Param input body RegisterMicroserviceRequest true "Register payload"
// @Success 201 "Ok."
// @Failure 400,403,500 {object} types.Error "Error."
// @Router /register/microservice [post]
func RegisterMicroservice(c *gin.Context, cfg *config.Configuration) {
	clientIp := c.ClientIP()

	if !utils.IsAllowedAddress(clientIp, cfg.MicroservicesIps) {
		c.JSON(http.StatusForbidden, types.Error{
			Error: "Address not allowed: " + clientIp,
		})
		return
	}

	var request RegisterMicroserviceRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, types.Error{
			Error: err.Error(),
		})
		return
	}

	exists, err := repository.UserExist(request.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.Error{
			Error: "Failed to check user",
		})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, types.Error{
			Error: "User already exist",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.Error{
			Error: "Failed to hash password",
		})
		return
	}

	user := models.User{
		Username:       request.Username,
		HashedPassword: hashedPassword,
		Role:           "MICROSERVICE",
	}

	err = repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.Error{
			Error: "Failed to register user",
		})
		return
	}

	c.Status(http.StatusCreated)
}

package handlers

import (
	"authentication/internal/database/repository"
	"authentication/internal/handlers/types"
	"authentication/internal/models"
	"authentication/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterUserRequest represents the payload required to register a new user.
type RegisterUserRequest struct {
	// Username of the user to be registered.
	// Required: true
	// Example: Zalimannard
	Username string `json:"username" binding:"required,alphanum,min=3,max=30"`

	// Password of the user to be registered.
	// Required: true
	// Example: querty01
	Password string `json:"password" binding:"required,alphanum,min=8,max=50"`
}

// @Summary Register a new user
// @Description Register a new user with the input payload
// @Accept  json
// @Produce  json
// @Param input body RegisterUserRequest true "Register payload"
// @Success 201 "Ok."
// @Failure 400,500 {object} types.Error "Error."
// @Router /register/user [post]
func Register(c *gin.Context) {
	var request RegisterUserRequest

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
		Role:           "USER",
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

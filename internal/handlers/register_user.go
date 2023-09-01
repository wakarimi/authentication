package handlers

import (
	"authentication/internal/database/repository"
	"authentication/internal/models"
	"authentication/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterRequest represents the payload required to register a new user.
type RegisterRequest struct {
	// Username of the user to be registered.
	// Required: true
	// Example: Zalimannard
	Username string `json:"username" binding:"required"`

	// Password of the user to be registered.
	// Required: true
	// Example: querty01
	Password string `json:"password" binding:"required"`
}

// Error represents a standard application error.
type Error struct {
	// The error message.
	// Example: User already exists
	Error string `json:"error" binding:"required"`
}

// @Summary Register a new user
// @Description Register a new user with the input payload
// @Accept  json
// @Produce  json
// @Param input body RegisterRequest true "Register payload"
// @Success 201 "Ok."
// @Failure 400,500 {object} Error "Error."
// @Router /register [post]
func RegisterUser(c *gin.Context) {
	var request RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, Error{
			Error: err.Error(),
		})
		return
	}

	exists, err := repository.UserExist(request.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error{
			Error: "Failed to check user",
		})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, Error{
			Error: "User already exist",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error{
			Error: "Failed to hash password",
		})
		return
	}

	user := models.User{
		Username:       request.Username,
		HashedPassword: hashedPassword,
	}

	err = repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error{
			Error: "Failed to register user",
		})
		return
	}

	c.Status(http.StatusCreated)
}

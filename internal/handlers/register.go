package handlers

import (
	"authentication/internal/database"
	"authentication/internal/database/repository"
	"authentication/internal/handlers/types"
	"authentication/internal/models"
	"authentication/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterAccountRequest represents the payload required to register a new account.
type RegisterAccountRequest struct {
	// Username of the account to be registered.
	// Required: true
	// Example: Zalimannard
	Username string `json:"username" binding:"required,alphanum,min=3,max=30"`

	// Password of the account to be registered.
	// Required: true
	// Example: qwerty01
	Password string `json:"password" binding:"required,alphanum,min=8,max=50"`
}

// Register
// @Summary Register a new account
// @Description Register a new account with the input payload. The first registered account will automatically be assigned the ADMIN role.
// @Accept  json
// @Produce  json
// @Param input body RegisterAccountRequest true "Register payload"
// @Success 201 "Ok."
// @Failure 400,500 {object} types.Error "Error."
// @Router /register [post]
func Register(c *gin.Context) {
	var request RegisterAccountRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, types.Error{
			Error: err.Error(),
		})
		return
	}

	exists, err := repository.AccountExist(request.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.Error{
			Error: "Failed to check account",
		})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, types.Error{
			Error: "Account already exist",
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

	count, err := repository.CountAccount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.Error{
			Error: "Failed to count accounts",
		})
		return
	}
	var role string
	if count == 0 {
		role = "ADMIN"
	} else {
		role = "USER"
	}

	account := models.Account{
		Username:       request.Username,
		HashedPassword: hashedPassword,
	}

	tx, err := database.Db.Beginx()
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.Error{
			Error: "Failed to start transaction",
		})
		return
	}

	accountId, err := repository.CreateAccountTx(tx, account)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, types.Error{
			Error: "Failed to register account",
		})
		return
	}

	err = repository.AssignRoleTx(tx, accountId, role)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, types.Error{
			Error: "Failed to assign role to account",
		})
		return
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.Error{
			Error: "Failed to commit transaction",
		})
		return
	}

	c.Status(http.StatusCreated)
}

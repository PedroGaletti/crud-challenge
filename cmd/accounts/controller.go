package accounts

import (
	"errors"
	"pismo/helper"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IAccountController: interface of Account controller
type IAccountController interface {
	Store(*gin.Context)
	Show(*gin.Context)
}

// AccountController: struct of Account controller
type AccountController struct {
	repository IAccountRepository
}

// NewAccountController: create a new Account controller
func NewAccountController(repository IAccountRepository) IAccountController {
	return &AccountController{repository}
}

// Store: store a new account in the database
func (c *AccountController) Store(ctx *gin.Context) {
	account := &Account{}

	if err := ctx.BindJSON(&account); err != nil {
		helper.BadRequestBodyReponse(ctx, err)
		return
	}

	if len(account.DocumentNumber) == 0 {
		helper.BadRequestBodyReponse(ctx, errors.New("you passed an empty document number"))
		return
	}

	if err := c.repository.Create(account); err != nil {
		helper.InternalServerErrorResponse(ctx, err)
		return
	}

	helper.OkResponse(ctx)
}

// Show: Get a specific account from database and return the information
func (c *AccountController) Show(ctx *gin.Context) {
	pId := ctx.Param("id")

	id, err := strconv.ParseInt(pId, 10, 64)
	if err != nil {
		helper.BadRequestBodyReponse(ctx, errors.New("you passed an invalid accountId query param"))
		return
	}

	account, err := c.repository.FindOne(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.NoContentResponse(ctx, err)
			return
		}

		helper.InternalServerErrorResponse(ctx, err)
		return
	}

	helper.OkDataResponse(ctx, account)
}

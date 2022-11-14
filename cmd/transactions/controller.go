package transactions

import (
	"errors"
	"pismo/helper"

	"github.com/gin-gonic/gin"
)

// ITransactionController: interface of Transaction controller
type ITransactionController interface {
	Store(*gin.Context)
}

// TransactionController: struct of Transaction controller
type TransactionController struct {
	repository ITransactionRepository
}

// NewTransactionController: create a new Transaction controller
func NewTransactionController(repository ITransactionRepository) ITransactionController {
	return &TransactionController{repository}
}

// Store: store a new transaction in the database
func (c *TransactionController) Store(ctx *gin.Context) {
	transaction := &Transaction{}

	if err := ctx.BindJSON(&transaction); err != nil {
		helper.BadRequestBodyReponse(ctx, err)
		return
	}

	if transaction.Amount == 0 {
		helper.BadRequestBodyReponse(ctx, errors.New("amount can't be zero"))
		return
	}

	if err := c.repository.Create(transaction); err != nil {
		helper.InternalServerErrorResponse(ctx, err)
		return
	}

	helper.OkResponse(ctx)
}

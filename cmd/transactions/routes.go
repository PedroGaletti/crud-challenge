package transactions

import (
	"github.com/gin-gonic/gin"
)

// Router : starting Transactions handler
func Router(r gin.IRoutes, transactionsController ITransactionController) {
	r.POST("", transactionsController.Store)
}

package accounts

import (
	"github.com/gin-gonic/gin"
)

// Router : starting Account handler
func Router(r gin.IRoutes, accountController IAccountController) {
	r.POST("", accountController.Store)
	r.GET("/:accountId", accountController.Show)
}

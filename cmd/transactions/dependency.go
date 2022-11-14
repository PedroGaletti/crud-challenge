package transactions

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	Repository ITransactionRepository
	Controller ITransactionController
)

// InjectDependency: create the injections of Transaction
func InjectDependency(router gin.IRoutes, db *gorm.DB) {
	// Repository
	Repository := NewTransactionRepository(db)

	// Controller
	Controller := NewTransactionController(Repository)

	// Routes
	Router(router, Controller)
}

package accounts

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	// Repository IAccountRepository
	Controller IAccountController
)

// InjectDependency: create the injections of Account
func InjectDependency(router gin.IRoutes, db *gorm.DB) {
	// Repository
	Repository := NewAccountRepository(db)

	// Controller
	Controller := NewAccountController(Repository)

	// Routes
	Router(router, Controller)
}

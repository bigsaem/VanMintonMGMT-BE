package routes

import (
	"github.com/gin-gonic/gin"
	// Import other necessary packages
)

// RegisterRoutes sets up all the routes for the application.
func RegisterRoutes(router *gin.Engine) {
	registerUserRoutes(router)
	registerTransactionRoutes(router)
	registerSessionRoutes(router)
	registerLedgerRoutes(router)
	// ... any other route registrations
}

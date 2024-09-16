package routes

import (
	"vanminton-be/internal/handlers"

	"github.com/gin-gonic/gin"
)

func registerLedgerRoutes(router *gin.Engine) {
	ledgerGroup := router.Group("/ledger")
	{
		ledgerGroup.POST("/", handlers.CreateLedgerEntry)
		// ... other ledger routes
	}
}

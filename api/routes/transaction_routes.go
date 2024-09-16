package routes

import (
	"vanminton-be/internal/handlers"

	"github.com/gin-gonic/gin"
)

func registerTransactionRoutes(router *gin.Engine) {
	transactionGroup := router.Group("/transactions")
	{
		transactionGroup.POST("/", handlers.CreateTransaction)
		// ... other transaction routes
	}
}

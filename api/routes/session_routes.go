package routes

import (
	"vanminton-be/internal/handlers"

	"github.com/gin-gonic/gin"
)

func registerSessionRoutes(router *gin.Engine) {
	sessionGroup := router.Group("/sessions")
	{
		sessionGroup.POST("/", handlers.CreateSession)
		// ... other session routes
	}
}

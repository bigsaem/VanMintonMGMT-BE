package routes

import (
	"vanminton-be/internal/handlers"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", handlers.CreateUser)
		userGroup.GET("/:id", handlers.GetUser)
		// ... other user routes like PUT, DELETE
	}
}

package router

import (
	"ms-ga-user/internal/api/handler"
	"ms-ga-user/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler, profileHandler *handler.ProfileHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CorrelationIDMiddleware())

	v1 := r.Group("/gymapi/v1")
	v1.Use(middleware.AuthMiddleware())

	users := v1.Group("/users")
	{
		users.GET("", middleware.RequirePermission("user:manage"), userHandler.ListUsers)
		users.POST("", middleware.RequirePermission("user:manage"), userHandler.CreateUser)
		users.GET("/search", middleware.RequirePermission("user:manage"), userHandler.SearchUsers)
		users.GET("/:id", userHandler.GetUserByID)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", middleware.RequirePermission("user:manage"), userHandler.DeactivateUser)

		users.GET("/:id/profile", profileHandler.GetProfile)
		users.PUT("/:id/profile", profileHandler.UpdateProfile)
	}

	return r
}

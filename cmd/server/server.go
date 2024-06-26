package server

import (
	"go_boilerplate/internal/api/auth"
	"go_boilerplate/internal/api/user"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	api := router.Group("/api")
	api.POST("/register", auth.Register)
	api.POST("/login", auth.Login)
	// api.Use(auth.AuthMiddleware)
	// api.GET("/users", user.GetUsers)
	api.POST("/user", user.CreateUser)
	// api.GET("/user/:id", user.GetUserById)
	// api.PUT("/user/:id", user.UpdateUser)
	api.DELETE("/user/:id", user.DeleteUser)
}

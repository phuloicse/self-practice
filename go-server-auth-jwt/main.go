package main

import (
	"jwt-app/handler"
	"jwt-app/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", handler.LoginHandler)

	protected := r.Group("/protected")
	protected.Use(middleware.JWTAuthMiddleware())
	protected.GET("/hello", handler.ProtectedHandler)

	r.Run(":8080")
}

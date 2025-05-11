package main

import (
	handler "jwt-app/handler"
	middleware "jwt-app/middleware"

	gin "github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	app.POST("/login", handler.LoginHandler)

	protected := app.Group("/protected")
	protected.Use(middleware.JWTAuthMiddleware())
	protected.GET("/hello", handler.ProtectedHandler)

	app.Run(":8080")
}

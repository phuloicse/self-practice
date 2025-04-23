package handler

import (
	service "jwt-app/service"

	utils "jwt-app/utils"

	model "jwt-app/model"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
)

func LoginHandler(requestContext *gin.Context) {
	var req model.User
	if err := requestContext.ShouldBindJSON(&req); err != nil {
		requestContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	log.Println("Request data:", req.Username, " ", req.Password)

	if err := service.Authenticate(req.Username, req.Password); err != nil {
		requestContext.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password"})
		return
	}

	log.Println("Generating JWT")
	token, err := utils.GenerateJWT(req.Username)
	if err != nil {
		requestContext.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	requestContext.JSON(http.StatusOK, gin.H{"token": token})
}

func ProtectedHandler(requestContext *gin.Context) {
	username := requestContext.GetString("username")
	requestContext.JSON(http.StatusOK, gin.H{"message": "Hello " + username + ", this is a protected route!"})
}

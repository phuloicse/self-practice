package config

import (
	"log"
	"os"
	"strconv"

	model "jwt-app/model"

	"github.com/joho/godotenv"
)

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetEnvAsInt(key string, defaultValue int) int {
	if valueStr := os.Getenv(key); valueStr != "" {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return defaultValue
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using defaults.")
	}

	JWTConfig = model.JWTInfo{
		SecretKey:   GetEnv("SECRET_KEY", "default_secret_key"),
		ExpiredTime: GetEnvAsInt("EXPIRED_TIME", 3600),
	}

	User = model.User{
		Username: GetEnv("LOGIN_USERNAME", "loi"),
		Password: GetEnv("LOGIN_PASSWORD", "12"),
	}

	log.Println("JWT Secret Key:", JWTConfig.SecretKey)
	log.Println("JWT Expired Time:", JWTConfig.ExpiredTime)
	log.Println("User Username:", User.Username)
	log.Println("User Password:", User.Password)
	log.Println("Environment variables loaded successfully.")	

	//  debug env in os
	// for _, e := range os.Environ() {
    // 	log.Println(e)
	// }

}

var JWTConfig model.JWTInfo
var User model.User

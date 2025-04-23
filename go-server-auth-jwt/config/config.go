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
		Username: GetEnv("USERNAME", "loi"),
		Password: GetEnv("PASSWORD", "12"),
	}

}

var JWTConfig model.JWTInfo
var User model.User

// func LoadConfig() {
// 	JWTConfig = model.JWTInfo{
// 		SecretKey:   GetEnv("SECRET_KEY", "default_secret_key"),
// 		ExpiredTime: GetEnvAsInt("EXPIRED_TIME", 3600),
// 	}

// 	User = model.User{
// 		Username: GetEnv("USERNAME", "loi"),
// 		Password: GetEnv("PASSWORD", "12"),
// 	}
// }

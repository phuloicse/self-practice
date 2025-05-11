package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"

	config "jwt-app/config"
)

var jwtKey = []byte(config.JWTConfig.SecretKey)

func GenerateJWT(username string) (string, error) {
	log.Println("GenerateJWT: Start generating token for user:", username)

	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
	}
	log.Printf("GenerateJWT: Claims created: %+v\n", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Println("GenerateJWT: Token object created with header SigningMethodHS256 and payload claims")

	signedToken, err := token.SignedString(jwtKey)
	log.Println("GenerateJWT: Token is signed with key ", jwtKey, " which is byte values of ", config.JWTConfig.SecretKey)
	if err != nil {
		log.Println("GenerateJWT: Error while signing token:", err)
		return "", err
	}

	log.Println("GenerateJWT: Result is ", signedToken)
	return signedToken, nil
}

func ValidateToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}

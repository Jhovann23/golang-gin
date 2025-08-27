package helper

import (
	"backend/config" 
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func GenerateToken(username string) string {

	expTime := time.Now().Add(60 * time.Minute)

	claims := &jwt.RegisteredClaims{
		Subject: username,
		ExpiresAt: jwt.NewNumericDate(expTime),

	}

	//membuat token baru 
	jwtToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)

	return jwtToken
}


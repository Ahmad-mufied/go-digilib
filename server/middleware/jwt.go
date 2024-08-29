package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTCustomClaims struct {
	UserID uint   `json:"userID"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, email, role string) (string, error) {
	claims := &JWTCustomClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as a response
	signedString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

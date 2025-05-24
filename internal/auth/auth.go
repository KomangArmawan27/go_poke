package auth

import (
	"errors"
	"log"
	"strconv"
	"time"

	"go-api/config"

	"github.com/golang-jwt/jwt/v5"
)

// Claims structure for JWT
type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT token
func GenerateToken(email, role string) (string, error) {
	jwtSecret := []byte(config.GetEnv("JWT_SECRET"))
	expiredInStr := config.GetEnv("JWT_EXPIRED_IN")

	// Convert to int (hours)
	expiredIn, err := strconv.Atoi(expiredInStr)
	if err != nil {
		log.Println("Invalid JWT_EXPIRED_IN format. Defaulting to 1 hour.")
		expiredIn = 1
	}

	expirationTime := time.Now().Add(time.Duration(expiredIn) * time.Hour) 
	claims := &Claims{
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken verifies a JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	jwtSecret := []byte(config.GetEnv("JWT_SECRET"))
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

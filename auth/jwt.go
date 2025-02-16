package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

// GenerateJWT generates a JWT token with the given secret and claims.
func GenerateJWT(secretKey string) (string, error) {
	// Define the claims
	claims := &jwt.MapClaims{
		"authorized": true,
		"user":       "test_user",
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	}

	// Create a new token object with signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Validation of token
func ValidateJWT(tokenStr string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid claims structure")
	}

	return &claims, nil
}

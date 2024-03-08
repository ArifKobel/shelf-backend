package auth_service

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(email string) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET"))
	claims := jwt.MapClaims{
		"email": email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetEmailFromToken(token string) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET"))
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}
	email := claims["email"].(string)
	return email, nil
}

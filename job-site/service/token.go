package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("0123456789")

// Helps to create a token
func (Service) CreateToken(username, email, role, roleID, userID string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"email":    email,
			"user_id":  userID,
			"role_id":  roleID,
			"role":     role,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Helps to verify the token is valid or not
func (Service) VerifyToken(tokenString string) (tokenData []byte, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return
	}

	if !token.Valid {
		return tokenData, fmt.Errorf("invalid token")
	}

	tokenData, err = json.Marshal(token.Claims)
	return
}

// Helps to create token without claims
func CreateTokenWithoutClaims(id string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  id,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

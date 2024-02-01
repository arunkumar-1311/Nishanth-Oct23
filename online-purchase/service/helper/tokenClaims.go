package helper

import (
	"encoding/json"
	"online-purchase/models"
)

// Helps to fetch the claims from the given JWT token
func Claims(token string, dest *models.Claims) error {
	claims, err := VerifyToken(token[7:])
	if err != nil {
		return err
	}

	if err := json.Unmarshal(claims, &dest); err != nil {
		return err
	}
	return nil
}

package helper

import (
	"blog_post/models"
	"encoding/json"
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

package helper

import (
	"blog_post/models"
	"encoding/json"
)

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

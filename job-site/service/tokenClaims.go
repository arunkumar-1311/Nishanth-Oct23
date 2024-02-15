package service

import (
	"job-post/models"
	"encoding/json"
)

// Helps to fetch the claims from the given JWT token
func (s Service)Claims(token string, dest *models.Claims) error {
	claims, err := s.VerifyToken(token[7:])
	if err != nil {
		return err
	}

	if err := json.Unmarshal(claims, &dest); err != nil {
		return err
	}
	return nil
}

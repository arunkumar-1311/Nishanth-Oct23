package service

import (
	"encoding/json"
	"to-do/models"
)

type Claims interface {
	Claims(token string, dest *models.Claims) error
}

// Helps to fetch the claims from the given JWT token
func (s Service) Claims(token string, dest *models.Claims) error {
	claims, err := s.VerifyToken(token)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(claims, &dest); err != nil {
		return err
	}
	return nil
}

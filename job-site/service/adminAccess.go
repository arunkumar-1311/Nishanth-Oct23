package service

import (
	"fmt"
	"job-post/models"
)

// Helps to check wheather the logined user is admin or not
func (s Service) AdminAccess(tokenString string) error {
	var claimsDetails models.Claims
	err := s.Claims(tokenString, &claimsDetails)
	if err != nil {
		return err
	}

	if claimsDetails.Role != "Admin" {
		return fmt.Errorf("admin can only access this feature")
	}
	return nil
}

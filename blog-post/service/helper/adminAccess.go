package helper

import (
	"blog_post/models"
	"fmt"
)

// Helps to check wheather the logined user is admin or not
func AdminAccess(tokenString string) error {
	var claimsDetails models.Claims
	err := Claims(tokenString, &claimsDetails)
	if err != nil {
		return err
	}

	if claimsDetails.RolesID != "AD1" {
		return fmt.Errorf("admin can only access this feature")
	}
	return nil
}

package helper

import (
	"blog_post/adaptor"
	"blog_post/models"
	"errors"
)

// It check wheather the emailid is already exist
func EmailAndNameValidation(user models.Users, db adaptor.Database) (result error) {

	userData, err := db.FindUserAndEmail(user.Email, user.Name)
	if err != nil {
		return err
	}

	if userData.Email != "" {
		if userData.Name == user.Name {
			return errors.New("user name already exist")
		}
		return errors.New("email id already exist")
	}

	return
}

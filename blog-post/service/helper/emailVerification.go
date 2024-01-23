package helper

import (
	"blog_post/models"
	"blog_post/repository"
	"errors"
)

// It check wheather the emailid is already exist
func EmailAndNameValidation(user models.Users) (result error) {

	userData, err := repository.FindUserAndEmail(user.Email, user.Name)
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

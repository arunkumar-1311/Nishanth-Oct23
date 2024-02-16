package service

import (
	"errors"
	"to-do/adaptor"
	"to-do/models"
)

type EmailAndNameValidation interface {
	EmailAndNameValidation(user models.Users, db adaptor.Database) (result error)
}

// It check wheather the emailid is already exist
func (Service) EmailAndNameValidation(user models.Users, db adaptor.Database) (result error) {

	userData, err := db.FindUserAndEmail(user.Email, user.UserName)
	if err != nil {
		return err
	}

	if userData.Email != "" {
		if userData.UserName == user.UserName {
			return errors.New("user name already exist")
		}
		return errors.New("email id already exist")
	}

	return
}

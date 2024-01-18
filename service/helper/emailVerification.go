package helper

import (
	"blog_post/repository"
)

func EmailValidation(email string) (result bool) {
	emailID, err := repository.FindEmail(email)
	if err != nil {
		return
	}
	if emailID != "" {
		return
	}
	return true
}

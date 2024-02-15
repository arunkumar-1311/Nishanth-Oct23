package service

import (
	"golang.org/x/crypto/bcrypt"
)

type Password interface {
	GenerateHash(password *string) error
	CompareHashPassword(password, hash string) bool
}

// Helps to generate the hash of given password
func (Service) GenerateHash(password *string) error {
	if *password == "" {
		return nil
	}
	byte, err := bcrypt.GenerateFromPassword([]byte(*password), 14)
	*password = string(byte)
	return err

}

// compare the given password and the hash
func (Service) CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

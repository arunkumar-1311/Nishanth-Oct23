package helper

import "golang.org/x/crypto/bcrypt"

// Helps to generate the hash of given password
func GenerateHash(password *string) error {
	byte, err := bcrypt.GenerateFromPassword([]byte(*password), 14)
	*password = string(byte)
	return err

}

// compare the given password and the hash
func CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

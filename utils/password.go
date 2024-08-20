package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func VerifyHashPassword(password, hashed_password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
	return err
}

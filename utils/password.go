package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password: %w", err)
	}

	return string(hashedPwd), nil
}

func VerifyPassword(hashedPwd, enteredPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(enteredPwd))
}

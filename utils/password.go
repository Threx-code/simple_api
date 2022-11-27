package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password, %s", err.Error())
	}

	return string(hashed), nil
}

func VerifyPassword(hasehPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasehPassword), []byte(password))
}

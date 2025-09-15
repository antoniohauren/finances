package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pw string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("error hashing the password %w", err)
	}

	return string(hashed), nil
}

func CheckPassword(pw string, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pw))
}

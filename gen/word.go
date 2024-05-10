package gen

import (
	"fmt"
	"crypto/rand"
	"math/big"
)

const (
	upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower   = "abcdefghijklmnopqrstuvwxyz"
	digits  = "0123456789"
	symbols = "!@#$%^&*"
)

func GeneratePassword(length int) (string, error) {
	if length <= 8 {
		return "", fmt.Errorf("Password length must be greater than 8 characters.")
	}

	allCharacters := upper + lower + digits + symbols

	var password string
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(allCharacters))))
		if err != nil {
			return "", err
		}
		
		char := allCharacters[index.Int64()]
		password += string(char)
	}

	return password, nil
}

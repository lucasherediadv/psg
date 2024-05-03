package gen

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower   = "abcdefghijklmnopqrstuvwxyz"
	digits  = "0123456789"
	symbols = "!@#$%^&*()_+{}[]|:;<>,.?/~"
)

func GeneratePassword(length int) (string, error) {
	if length <= 6 {
		return "", fmt.Errorf("Invalid password length. Password must be 6 character or more.")
	}

	allCharacters := upper + lower + digits + symbols

	rand.Seed(time.Now().UnixNano())

	var password string
	for i := 0; i < length; i++ {
		char := allCharacters[rand.Intn(len(allCharacters))]

		password += string(char)
	}

	return password, nil
}

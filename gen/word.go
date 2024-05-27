package gen

import (
	"fmt"
	"crypto/rand"
	"math/big"
)

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars  = "0123456789"
	symbolChars = "!@#$%^&*()-_=+{}[]|:;<>,.?/~`"
)

func GeneratePassword(length int, upper bool, numbers bool, symbols bool,) (string, error) {
	if length <= 8 {
		return "", fmt.Errorf("Password length must be greater than 8 characters.")
	}

	allCharacters := lowerChars

	if upper {
		allCharacters += upperChars 
	}
	if numbers {
		allCharacters += numberChars
	}
	if symbols {
		allCharacters += symbolChars
	}

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

package gen

import (
	"crypto/rand"
	_ "embed"
	"fmt"
	"math/big"
	"strings"
)

//go:embed "wordlist/eff_large_wordlist.txt"
var wordList string

var words = strings.Split(wordList, "\n")

func GeneratePassphrase(length int, separator string, capitalize bool) (string, error) {
	if len(words) == 0 {
		return "", fmt.Errorf("Word list is empty")
	}

	if length <= 0 || length > 100 {
		return "", fmt.Errorf("Passphrase length must be a positive integer and not exceed 100")
	}

	if separator == "" {
		return "", fmt.Errorf("Separator must not be empty")
	}

	passphrase := make([]string, length)
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
		if err != nil {
			return "", fmt.Errorf("Failed to generate random index: %w", err)
		}
		word := words[randomIndex.Int64()]

		if capitalize {
			word = strings.Title(word)
		}

		passphrase[i] = word
	}

	return strings.Join(passphrase, separator), nil
}

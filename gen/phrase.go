package gen

import (
	"crypto/rand"
	"math/big"
	"os"
	"strings"
)

var wordList = "gen/wordlist/eff_large_wordlist.txt"

func GeneratePassphrase(length int, separator string, capitalize bool) (string, error) {
	content, err := os.ReadFile(wordList)
	if err != nil {
		return "", err
	}

	words := strings.Split(string(content), "\n")

	var passphrase []string
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
		if err != nil {
			return "", err
		}
		word := words[randomIndex.Int64()]

		passphrase = append(passphrase, word)
	}

	// Capitalize every word in the passphrase
	if capitalize {
		for i, word := range passphrase {
			passphrase[i] = strings.Title(word)
		}
	}

	return strings.Join(passphrase, separator), nil
}

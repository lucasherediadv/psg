package gen

import (
	"fmt"
	"crypto/rand"
	"math/big"
	"os"
	"strings"
)

var WordLists = map[string][]string {
	"en": {"gen/wordlists/wordlist_en.txt", "English"},
	"es": {"gen/wordlists/wordlist_es.txt", "Español"},
}

func GeneratePassphrase(length int, separator string, capitalize bool, wordList string) (string, error) {
	wordListFile, ok := WordLists[wordList]
	if !ok {
		return "", fmt.Errorf("Not a valid word list: %s", wordList)
	}

	content, err := os.ReadFile(wordListFile[0])
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

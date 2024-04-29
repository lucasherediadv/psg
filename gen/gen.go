package gen

import (
	"crypto/rand"
	"math/big"
	"os"
	"strings"
)

func ReadWords() ([]string, error) {
	wordList := "wordlist/eff_large_wordlist.txt"
	content, err := os.ReadFile(wordList)
	if err != nil {
		return nil, err
	}

	words := strings.Split(string(content), "\n")
	return words, nil
}

func PassGen(words []string, wordCount int, sep string) (string, error) {
	var passphrase []string
	for i := 0; i < wordCount; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
		if err != nil {
			return "", err
		}
		passphrase = append(passphrase, words[randomIndex.Int64()])
	}

	return strings.Join(passphrase, sep), nil
}

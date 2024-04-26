package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func generate(wordlist string, wordCount int, sep string) (string, error) {
	// Read words from the specified file
	content, err := os.ReadFile(wordlist)
	if err != nil {
		return "", err
	}

	// Split the content into individual words
	words := strings.Split(string(content), "\n")

	// Generate the passphrase
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

func main() {
	wordlist := "wordlist/eff_large_wordlist.txt"
	wordCount := 8
	separator := "-"

	passphrase, err := generate(wordlist, wordCount, separator)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(passphrase)
}

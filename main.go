package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
)

var (
	wordCount = flag.Int("count", 8, "Number of words in the passphrase")
	separator = flag.String("sep", "-", "Separator between words")
)

func readWords(wordlist string) ([]string, error) {
	content, err := os.ReadFile(wordlist)
	if err != nil {
		return nil, err
	}

	words := strings.Split(string(content), "\n")
	return words, nil
}

func passGen(words []string, wordCount int, sep string) (string, error) {
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
	flag.Parse()

	wordlist := "wordlist/eff_large_wordlist.txt"

	words, err := readWords(wordlist)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	passphrase, err := passGen(words, *wordCount, *separator)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(passphrase)
}

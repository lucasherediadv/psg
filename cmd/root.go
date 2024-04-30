package cmd

import (
	"flag"
	"fmt"

	"github.com/lucasherediadv/passgen/gen"
)

var (
	wordCount = flag.Int("count", 8, "Number of words in the passphrase")
	separator = flag.String("sep", "-", "Separator between words")
)

func Execute() {
	flag.Parse()

	words, err := gen.ReadWords()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	passphrase, err := gen.PassGen(words, *wordCount, *separator)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(passphrase)
}

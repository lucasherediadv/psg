package cmd

import (
	"flag"
	"fmt"

	"github.com/lucasherediadv/passgen/gen"
)

var (
	wordCount = flag.Int("count", 8, "Number of words in the passphrase")
	separator = flag.String("sep", "-", "Separator between words")
    showUsage = flag.Bool("help", false, "Show usage information")
)

func Execute() {
	flag.Parse()

    if *showUsage {
        fmt.Println("Generate secure passphrases using the diceware method\n")
        fmt.Println("Usage: \n  passgen [options]\n")
        fmt.Println("Options:")
        flag.PrintDefaults()
        return
    }

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

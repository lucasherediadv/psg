package cmd

import (
	"flag"
	"fmt"

	"github.com/lucasherediadv/passgen/gen"
)

var (
	wordCount = flag.Int("c", 8, "Number of words in the passphrase")
	separator = flag.String("s", "-", "Separator between words")
    showUsage = flag.Bool("h", false, "Show usage information")
)

func Execute() {
	flag.Parse()

    if *showUsage {
        fmt.Println("Usage: passgen [options]")
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

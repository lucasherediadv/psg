package cmd

import (
	"fmt"

	"github.com/lucasherediadv/passgen/gen"
	"github.com/spf13/cobra"
)

var phraseCmd = &cobra.Command{
	Use:   "phrase",
	Short: "Generate a passphrase",
	Long:  `Generate a passphrase with the default values`,
	Run: func(cmd *cobra.Command, args []string) {
		wordList, err := cmd.Flags().GetString("word-list")
		if err != nil {
			fmt.Println("Error getting word list: ", err)
			return
		}

		length, err := cmd.Flags().GetInt("length")
		if err != nil {
			fmt.Println("Error getting length: ", err)
			return
		}

		separator, err := cmd.Flags().GetString("separator")
		if err != nil {
			fmt.Println("Error getting separator: ", err)
			return
		}

		randomNumber, err := cmd.Flags().GetBool("random-number")
		if err != nil {
			fmt.Println("Error getting randomNumber: ", err)
			return
		}

		capitalize, err := cmd.Flags().GetBool("capitalize")
		if err != nil {
			fmt.Println("Error getting capitalize: ", err)
			return
		}

		passphrase, err := gen.GeneratePassphrase(length, separator, capitalize, randomNumber, wordList)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(passphrase)
	},
}

func init() {
	rootCmd.AddCommand(phraseCmd)

	phraseCmd.Flags().StringP("word-list", "w", "en", "Language from the wordlist")
	phraseCmd.Flags().IntP("length", "l", 6, "Number of words in the passphrase")
	phraseCmd.Flags().StringP("separator", "s", "-", "Separator character between words in the passphrase")
	phraseCmd.Flags().BoolP("random-number", "r", false, "Add a random number to the end of each word in the passphrase")
	phraseCmd.Flags().BoolP("capitalize", "c", false, "Capitalize every word in the passphrase")
}

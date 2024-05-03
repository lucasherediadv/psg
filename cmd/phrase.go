package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/lucasherediadv/passgen/gen"
)

var phraseCmd = &cobra.Command{
	Use:   "phrase",
	Short: "Generate a passphrase",
	Long:  `Generate a passphrase with the default values`,
	Run: func(cmd *cobra.Command, args []string) {
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

		randomNumber, err := cmd.Flags().GetBool("randomNumber")
		if err != nil {
			fmt.Println("Error getting randomNumber: ", err)
			return
		}

		capitalize, err := cmd.Flags().GetBool("capitalize")
		if err != nil {
			fmt.Println("Error getting capitalize: ", err)
			return
		}

		passphrase, err := gen.GeneratePassphrase(length, separator, capitalize, randomNumber)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(passphrase)
	},
}

func init() {
	rootCmd.AddCommand(phraseCmd)

	phraseCmd.Flags().IntP("length", "l", 6, "Number of words in the passphrase")
	phraseCmd.Flags().StringP("separator", "s", "-", "Separator character between words in the passphrase")
	phraseCmd.Flags().BoolP("randomNumber", "r", false, "Add a random number to the end of each word in the passphrase")
	phraseCmd.Flags().BoolP("capitalize", "c", false, "Capitalize every word in the passphrase")
}

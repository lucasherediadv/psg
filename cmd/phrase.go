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
		length, err := cmd.Flags().GetInt("length")
		if err != nil {
			fmt.Println("Error getting length: %v", err)
			return
		}

		separator, err := cmd.Flags().GetString("separator")
		if err != nil {
			fmt.Println("Error getting separator: %v", err)
			return
		}

		capitalize, err := cmd.Flags().GetBool("capitalize")
		if err != nil {
			fmt.Println("Error getting capitalize: %v", err)
			return
		}

		passphrase, err := gen.GeneratePassphrase(length, separator, capitalize)
		if err != nil {
			fmt.Println("Error generating passphrase: %v", err)
			return
		}

		fmt.Println(passphrase)
	},
}

func init() {
	rootCmd.AddCommand(phraseCmd)

	phraseCmd.Flags().IntP("length", "l", 8, "number of words in the passphrase")
	phraseCmd.Flags().StringP("separator", "s", "-", "separator character between words in the passphrase")
	phraseCmd.Flags().BoolP("capitalize", "c", false, "capitalize every word in the passphrase")
}

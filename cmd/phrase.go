package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/lucasherediadv/psg/gen"
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

		passphrase, err := gen.GeneratePassphrase(length, separator)
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
}

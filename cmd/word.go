package cmd

import (
	"fmt"

	"github.com/lucasherediadv/passgen/gen"
	"github.com/spf13/cobra"
)

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "Generate a password",
	Long:  `Generate a password with the default values`,
	Run: func(cmd *cobra.Command, args []string) {
		length, err := cmd.Flags().GetInt("length")
		if err != nil {
			fmt.Println("Error getting length: ", err)
			return
		}

		password, err := gen.GeneratePassword(length)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Println(password)
	},
}

func init() {
	rootCmd.AddCommand(wordCmd)

	wordCmd.Flags().IntP("length", "l", 20, "Number of characters in the password")
}

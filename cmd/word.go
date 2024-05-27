package cmd

import (
	"fmt"

	"github.com/lucasherediadv/psg/gen"
	"github.com/spf13/cobra"
)

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "Generate a password",
	Long:  `Generate a password with the default values`,
	Run: func(cmd *cobra.Command, args []string) {
		length, err := cmd.Flags().GetInt("length")
		if err != nil {
			fmt.Println("Error getting length: %v", err)
			return
		}

		entropy, err := cmd.Flags().GetBool("entropy")
		if err != nil {
			fmt.Println("Error getting entropy: %v", err)
			return
		}

		password, err := gen.GeneratePassword(length)
		if err != nil {
			fmt.Println("Error generating the password: %v", err)
			return
		}

		fmt.Println(password)
		if entropy {
			passwordEntropy := gen.CalculateEntropy(password)
			fmt.Printf("Entropy of the generated password: %.0f bits\n", passwordEntropy)
		}
	},
}

func init() {
	rootCmd.AddCommand(wordCmd)

	wordCmd.Flags().IntP("length", "l", 20, "number of characters in the password")
	wordCmd.Flags().BoolP("entropy", "e", false, "calculate password entropy")
}

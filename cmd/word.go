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

		upper, err := cmd.Flags().GetBool("upper")
		if err != nil {
			fmt.Println("Error getting upper: %v", err)
			return
		}

		numbers, err := cmd.Flags().GetBool("numbers")
		if err != nil {
			fmt.Println("Error getting numbers: %v", err)
			return
		}

		symbols, err := cmd.Flags().GetBool("symbols")
		if err != nil {
			fmt.Println("Error getting symbols: %v", err)
			return
		}

		entropy, err := cmd.Flags().GetBool("entropy")
		if err != nil {
			fmt.Println("Error getting entropy: %v", err)
			return
		}

		password, err := gen.GeneratePassword(length, upper, numbers, symbols)
		if err != nil {
			fmt.Println("Error generating the password: %v", err)
			return
		}

		fmt.Println(password)
		if entropy {
			passwordEntropy := gen.CalculateEntropy(password)
			fmt.Printf("Entropy of the generated password: %f bits\n", passwordEntropy)
		}
	},
}

func init() {
	rootCmd.AddCommand(wordCmd)

	wordCmd.Flags().IntP("length", "l", 20, "number of characters in the password")
	wordCmd.Flags().BoolP("entropy", "e", false, "calculate password entropy")
	wordCmd.Flags().BoolP("symbols", "s", false, "Include symbols in the password")
	wordCmd.Flags().BoolP("upper", "u", false, "Include upper case in the password")
	wordCmd.Flags().BoolP("numbers", "n", false, "Include numbers in the password")

}

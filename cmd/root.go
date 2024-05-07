package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "passgen",
	Short: "Generate secure passphrases",
	Long:  `Passphrase generator inspired by the diceware method to generate secure and memorable passphrases.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

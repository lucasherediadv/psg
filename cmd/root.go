package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "psg",
	Short:   "Passphrase and password generator",
	Long:    `Generates memorable passphrases and cryptographically secure passwords`,
	Version: "1.2.4",
}

func Execute() error {
	// Disable default completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Execute the root command
	err := rootCmd.Execute()
	if err != nil {
		return fmt.Errorf("Error executing root command: %w", err)
	}

	return nil
}

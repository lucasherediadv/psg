package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var rootCmd = &cobra.Command{
	Use:     "psg",
	Short:   "Passphrase and password generator",
	Long:    `Generates memorable passphrases and cryptographically secure passwords`,
	Version: "1.2.1",
}

func Execute() error {
	// Disable default completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Remove timestamp in generated documentation
	rootCmd.DisableAutoGenTag = true

	// Generate markdown documentation
	err := doc.GenMarkdownTree(rootCmd, "./doc")
	if err != nil {
		return fmt.Errorf("Error generating markdown documentation: %w", err)
	}

	// Execute the root command
	err = rootCmd.Execute()
	if err != nil {
		return fmt.Errorf("Error executing root command: %w", err)
	}

	return nil
}

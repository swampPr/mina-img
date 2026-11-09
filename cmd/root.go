// Package cmd provides cmd  ->  Provides the root command "mina-img"
package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mina-img",
	Short: "A very barebones image format focused on tiny file sizes",
	RunE:  execRoot,
}

func execRoot(cmd *cobra.Command, _ []string) error {
	_ = cmd.Help()
	return errors.New("you must select a subcommand. view subcommands with 'mina-img --help'")
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

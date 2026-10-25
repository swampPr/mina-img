// Package cmd provides cmd  ->  Provides the root command "mina-img"
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mina-img",
	Short: "A very barebones image format focused on tiny file sizes",
	Run:   execRoot,
}

func execRoot(cmd *cobra.Command, _ []string) {
	_ = cmd.Help()
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

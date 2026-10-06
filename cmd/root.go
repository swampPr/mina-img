// Package cmd provides cmd  ->  Provides the root command "mina-img"
package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mina-img",
	Short: "A very barebones image format focused on tiny file sizes",
	RunE:  execRoot,
}

func execRoot(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.New("you must provide the path of the source image and the output path of the destination image")
	}

	inPath, err := filepath.Abs(args[0])
	if err != nil {
		return err
	}

	outPath, err := filepath.Abs(args[1])
	if err != nil {
		return err
	}

	fmt.Println(inPath, outPath)

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

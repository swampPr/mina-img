package cmd

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/swampPr/mina-img/internal/mina"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:     "convert",
	Short:   "Convert a MINA image file to BMP, WEBP, PNG or JPEG",
	Example: "mina-img convert /path/to/source.mina /path/to/out.bmp",
	RunE:    execConvert,
}

func execConvert(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.New("you must provide a source path and an out path")
	}

	opts := &mina.Opts{
		Op:      mina.ConvertOp,
		InPath:  args[0],
		OutPath: args[1],
	}

	return mina.Handle(opts)
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

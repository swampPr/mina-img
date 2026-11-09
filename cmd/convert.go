package cmd

import (
	"errors"
	"strings"

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

func execConvert(_ *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.New("you must provide a source path and an out path")
	}

	opts := &mina.Opts{
		Op:      mina.ConvertOp,
		InPath:  strings.TrimSpace(args[0]),
		OutPath: strings.TrimSpace(args[1]),
	}

	return mina.Handle(opts)
}

func init() {
	rootCmd.AddCommand(convertCmd)
}

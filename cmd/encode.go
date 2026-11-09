package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"

	"github.com/swampPr/mina-img/internal/mina"
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:     "encode",
	Short:   "Encode an image (BMP, WEBP, JPEG, PNG) into MINA",
	Example: "mina-img encode /path/to/source.bmp /path/to/out.mina",
	RunE:    execEncode,
}

func execEncode(_ *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.New("you must provide a valid source path and destination path")
	}

	opts := &mina.Opts{
		InPath:  strings.TrimSpace(args[0]),
		OutPath: strings.TrimSpace(args[1]),
		Op:      mina.EncodeOp,
	}

	return mina.Handle(opts)
}

func init() {
	rootCmd.AddCommand(encodeCmd)
}

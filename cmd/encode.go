package cmd

import (
	"errors"

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
		return errors.New("you must provide a valid source path and outpath")
	}

	opts := &mina.Opts{
		InPath:  args[0],
		OutPath: args[1],
		Op:      mina.EncodeOp,
	}

	return mina.Handle(opts)
}

func init() {
	rootCmd.AddCommand(encodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

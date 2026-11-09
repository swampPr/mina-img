package mina

import (
	"errors"
	"os"
	"strings"

	"github.com/swampPr/mina-img/internal/utils"
)

// Opts struct  ->  The options used for any operation passed to the handler
type Opts struct {
	Op      Operation
	InPath  string
	OutPath string
}

type Operation = string

const (
	EncodeOp  Operation = "Encode"
	DecodeOp  Operation = "Decode"
	ConvertOp Operation = "Convert"
)

type HandlerFunc = func(o *Opts) error

var errBadOutPath = errors.New("bad destination path. outpath format: /path/to/out.format")

var HandlersRegistry = map[Operation]HandlerFunc{
	"Encode":  handleEncoding,
	"Convert": handleConvertion,
}

// Handle function  ->  Orchestrates all operations such as decoding, encoding or conversion
func Handle(o *Opts) error {
	if o.InPath == "" || o.OutPath == "" {
		return errors.New("source path and destination path are empty")
	}

	handler, ok := HandlersRegistry[o.Op]
	if !ok {
		return errors.New("invalid operation")
	}

	return handler(o)
}

func handleEncoding(o *Opts) error {
	sourceImg, err := utils.LoadImage(o.InPath)
	if err != nil {
		return err
	}

	newFile, err := os.Create(o.OutPath)
	if err != nil {
		return err
	}

	defer func() {
		_ = newFile.Close()
	}()

	return encode(newFile, sourceImg)
}

func handleConvertion(o *Opts) error {
	sourceFile, err := utils.LoadFile(o.InPath)
	if err != nil {
		return err
	}

	sourceImg, err := Decode(sourceFile)
	if err != nil {
		return err
	}

	_, format, _ := strings.CutLast(o.OutPath, ".")
	if len(format) < 1 {
		return errBadOutPath
	}

	newFile, err := os.Create(o.OutPath)
	if err != nil {
		return err
	}

	defer func() {
		_ = newFile.Close()
	}()

	return convert(newFile, sourceImg, format)
}

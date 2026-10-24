package mina

import (
	"errors"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"golang.org/x/image/bmp"

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

	return Encode(newFile, sourceImg)
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

	newFile, err := os.Create(o.OutPath)
	if err != nil {
		return err
	}

	defer func() {
		_ = newFile.Close()
	}()

	switch format {
	case "png":
		err = png.Encode(newFile, sourceImg)
		if err != nil {
			return err
		}
	case "jpeg":
		err = jpeg.Encode(newFile, sourceImg, nil)
		if err != nil {
			return err
		}
	case "bmp":
		err = bmp.Encode(newFile, sourceImg)
		if err != nil {
			return err
		}
	case "jpg":
		err = jpeg.Encode(newFile, sourceImg, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

package mina

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/deepteams/webp"
	"golang.org/x/image/bmp"
)

func convert(dst io.Writer, src image.Image, format string) error {
	switch format {
	case "png":
		err := png.Encode(dst, src)
		if err != nil {
			return err
		}
	case "jpeg":
		err := jpeg.Encode(dst, src, nil)
		if err != nil {
			return err
		}
	case "bmp":
		err := bmp.Encode(dst, src)
		if err != nil {
			return err
		}
	case "webp":
		err := webp.Encode(dst, src, nil)
		if err != nil {
			return err
		}
	case "jpg":
		err := jpeg.Encode(dst, src, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

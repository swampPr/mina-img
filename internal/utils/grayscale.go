package utils

import (
	"image"
	"image/color"
)

func Grayscale(i image.Image) *image.RGBA {
	bounds := i.Bounds()
	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := i.At(x, y).RGBA()

			r8, g8, b8 := byte(r>>8), byte(g>>8), byte(b>>8)

			gray := 0.299*float64(r8) + 0.587*float64(g8) + 0.114*float64(b8)

			dst.Set(x, y, color.Gray{Y: byte(gray)})

		}
	}

	return dst
}

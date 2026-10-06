package mina

import (
	"image"
	"image/color"
	"math"
)

type MinaImg struct {
	PixMap   []byte
	PixSize  int
	Height   int
	Width    int
	Metadata string
}

var luminanceColorModel color.Model = color.ModelFunc(luminanceModelConvert)

func luminanceModelConvert(c color.Color) color.Color {
	if gray, ok := c.(color.Gray); ok {
		nibble := math.Round((float64(gray.Y) / 255.0) * 10.0)
		b := math.Round(((nibble / 10.0) * 255.0))
		return color.Gray{Y: byte(b)}
	}

	gray := color.GrayModel.Convert(c).(color.Gray)

	nibble := math.Round((float64(gray.Y) / 255.0) * 10.0)

	b := math.Round(((nibble / 10.0) * 255.0))

	return color.Gray{Y: byte(b)}
}

func (i *MinaImg) ColorModel() color.Model {
	return luminanceColorModel
}

func (i *MinaImg) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (i *MinaImg) At(x, y int) color.Color {
	if !(image.Point{X: x, Y: y}.In(i.Bounds())) {
		return color.Gray{Y: 0}
	}

	stride := (i.Width + 1) / 2

	yi := y - i.Bounds().Min.Y
	xi := x - i.Bounds().Min.X

	index := yi*stride + (xi / 2)

	b := i.PixMap[index]

	var val byte
	if xi%2 == 0 {
		val = b >> 4
	} else {
		val = b & 0x0F
	}

	grayVal := math.Round((float64(val) / 10.0) * 255.0)

	return color.Gray{Y: uint8(grayVal)}
}

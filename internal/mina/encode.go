package mina

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"math"

	"github.com/swampPr/mina-img/internal/utils"
)

func Encode(w io.Writer, i image.Image) error {
	grayscaled := utils.Grayscale(i)

	img := createImage(grayscaled)

	writer := bufio.NewWriter(w)

	_, err := writer.WriteString(img.Metadata)
	if err != nil {
		return err
	}

	_, err = writer.Write(img.PixMap)
	if err != nil {
		return err
	}

	return writer.Flush()
}

func createImage(i image.Image) *MinaImg {
	bounds := i.Bounds()
	mapLength := (i.Bounds().Dx() / 2) * i.Bounds().Dy()
	pixMap := make([]byte, 0, mapLength)

	width := i.Bounds().Dx()
	height := i.Bounds().Dy()

	metadata := fmt.Sprintf("MIN-V1.1\n%d\n%d\r\n\r\n", width, height)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X-1; x++ {
			gray1, _, _, _ := i.At(x, y).RGBA()
			gray2, _, _, _ := i.At(x+1, y).RGBA()

			converted1 := math.Round(float64((gray1 / 255.0) * 10.0))
			converted2 := math.Round(float64((gray2 / 255.0) * 10.0))

			packed := (byte(converted1) << 4) | (byte(converted2) & 15)

			pixMap = append(pixMap, packed)

		}
	}

	return &MinaImg{
		PixMap:   pixMap,
		PixSize:  mapLength,
		Metadata: metadata,
		Width:    width,
		Height:   height,
	}
}

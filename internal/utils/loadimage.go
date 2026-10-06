package utils

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	_ "github.com/deepteams/webp"
	_ "golang.org/x/image/bmp"
)

// TODO: finish
func LoadImage(path string) (image.Image, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	src, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return src, nil
}

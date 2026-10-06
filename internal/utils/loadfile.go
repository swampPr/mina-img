package utils

import (
	"errors"
	"os"
)

func LoadFile(path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, errors.New("this file does not exist")
		}
		if errors.Is(err, os.ErrPermission) {
			return nil, errors.New("no permission to open this file")
		}

		return nil, err
	}

	return f, nil
}

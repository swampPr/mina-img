package mina

import (
	"bufio"
	"bytes"
	"errors"
	"image"
	"io"
	"log"
	"strconv"
	"strings"
)

// ErrMalformed  -> error for a malformed file
var ErrMalformed = errors.New("malformed headers")

// Decode function  ->  Decodes the returned reader and returns an image or error
func Decode(src io.Reader) (image.Image, error) {
	reader, headers, metadata, err := readHeaders(src)
	if err != nil {
		return nil, err
	}

	mapLength := ((headers.Width + 1) / 2) * headers.Height
	pixMap := make([]byte, mapLength)

	_, err = io.ReadFull(reader, pixMap)
	if err != nil {
		return nil, err
	}

	img := &MinaImg{
		PixMap:   pixMap,
		Height:   headers.Height,
		Width:    headers.Width,
		PixSize:  mapLength,
		Metadata: string(metadata),
	}

	return img, nil
}

func readHeaders(r io.Reader) (reader *bufio.Reader, headers MinaHeaders, metadata []byte, err error) {
	reader = bufio.NewReader(r)
	headerBuffer := make([]byte, 0, 20)

	delimiter := []byte("\r\n\r\n")

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) || errors.Is(err, bufio.ErrBufferFull) {
				return nil, MinaHeaders{}, nil, ErrMalformed
			}

			log.Println(err)
			return nil, MinaHeaders{}, nil, err
		}

		headerBuffer = append(headerBuffer, b)

		if before, ok := bytes.CutSuffix(headerBuffer, delimiter); ok {
			metadata = headerBuffer
			headerBuffer = before
			break
		}
	}

	header := strings.Split(string(headerBuffer), "\n")

	width, err := strconv.Atoi(header[1])
	if err != nil {
		return nil, MinaHeaders{}, nil, ErrMalformed
	}

	height, err := strconv.Atoi(header[2])
	if err != nil {
		return nil, MinaHeaders{}, nil, ErrMalformed
	}

	headers = MinaHeaders{
		Version: header[0],
		Width:   width,
		Height:  height,
	}

	return reader, headers, metadata, nil
}

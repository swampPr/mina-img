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

var ErrMalformed = errors.New("malformed headers")

func Decode(src io.Reader) (image.Image, error) {
	reader, width, height, metadata, err := readHeaders(src)
	if err != nil {
		return nil, err
	}

	mapLength := ((width / 2) + 1) * height
	pixMap := make([]byte, mapLength)

	_, err = io.ReadFull(reader, pixMap)
	if err != nil {
		return nil, err
	}

	img := &MinaImg{
		PixMap:   pixMap,
		Height:   height,
		Width:    width,
		PixSize:  mapLength,
		Metadata: string(metadata),
	}

	return img, nil
}

func readHeaders(r io.Reader) (reader *bufio.Reader, width, height int, metadata []byte, err error) {
	reader = bufio.NewReader(r)
	headerBuffer := make([]byte, 0, 20)

	delimiter := []byte("\r\n\r\n")

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) || errors.Is(err, bufio.ErrBufferFull) {
				return nil, 0, 0, nil, ErrMalformed
			}

			log.Println(err)
			return nil, 0, 0, nil, err
		}

		headerBuffer = append(headerBuffer, b)

		if before, ok := bytes.CutSuffix(headerBuffer, delimiter); ok {
			metadata = headerBuffer
			headerBuffer = before
			break
		}
	}

	header := strings.Split(string(headerBuffer), "\n")

	if width, err = strconv.Atoi(header[1]); err != nil {
		return nil, 0, 0, nil, ErrMalformed
	}

	if height, err = strconv.Atoi(header[2]); err != nil {
		return nil, 0, 0, nil, ErrMalformed
	}

	return reader, width, height, metadata, nil
}

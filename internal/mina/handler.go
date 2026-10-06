package mina

import "errors"

// Opts struct  ->  The options used for any operation passed to the handler
type Opts struct {
	Decode  bool
	Encode  bool
	InPath  string
	OutPath string
}

// Handle function  ->  Handles all operations such as decoding or encoding or conversion
func Handle(o *Opts) error {
	if o.InPath == "" || o.OutPath == "" {
		return errors.New("source path and destination path are empty")
	}

	return nil
}

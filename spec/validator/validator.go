package validator

import (
	"errors"
	"io"

	"github.com/gabriel-vasile/mimetype"
)

func Validate(r io.Reader, t string) (bool, error) {
	ftype, err := mimetype.DetectReader(r)
	if err != nil {
		return false, errors.New("can not detect file type")
	}

	return ftype.String() == t, nil
}

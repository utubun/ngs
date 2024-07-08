package validator

import (
	"errors"
	"fmt"
	"io"

	"github.com/gabriel-vasile/mimetype"
)

func Validate(r io.Reader, t string) (bool, error) {
	ftype, err := mimetype.DetectReader(r)
	if err != nil {
		return false, errors.New("can not detect file type")
	}

	fmt.Println("Validated!")
	return ftype.String() == t, nil
}

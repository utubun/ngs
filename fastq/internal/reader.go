package internal

import (
	"bufio"
	"fmt"
	"io"
)

type Reader struct {
	scanner *bufio.Scanner
}

func NewReader(r io.Reader) (*Reader, error) {
	scanner := bufio.NewScanner(r)
	return &Reader{scanner: scanner}, nil
}

func (r *Reader) Read(b []byte) (int, error) {
	// record identifier
	var p string
	// n holds number of bytes written
	var n int
	// address of the sbstring
	var address int8
	// tem holds temporary byte array
	temp := make([]byte, 3)

	for i := 0; i < 4; i++ {
		ok := r.scanner.Scan()
		if !ok {
			return 0, io.EOF
		}
		text := r.scanner.Text()
		kind := IdentifyReadLine(text)
		switch kind {
		case "header":
			if p != "" {
				return 0, fmt.Errorf("invalid format: expected header record, but found %s", text)
			}
			header := r.scanner.Bytes()
			address = int8(len(temp) + len(header))
			temp[1] = byte(address)
			temp = append(temp, r.scanner.Bytes()...)
			p = "header"
			continue
		case "dna":
			if p != "header" {
				return 0, fmt.Errorf("invalid format: expected dna record, but found %s", text)
			}
			p = "dna"
			dna := r.scanner.Bytes()
			address = int8(int(temp[1]) + len(dna))
			temp[2] = byte(address)
			temp = append(temp, dna...)
			continue
		case "util":
			if p != "dna" {
				return 0, fmt.Errorf("invalid format: expected utility record, but found %s", text)
			}
			p = "util"
			continue
		case "quality":
			if p != "util" {
				return 0, fmt.Errorf("invalid format: expected quality record, but found %s", text)
			}
			temp[0] = byte(1)
			temp = append(temp, r.scanner.Bytes()...)
		default:
			return 0, fmt.Errorf("invalid format: unknown record type %s", text)
		}
	}
	n += copy(b, temp)
	return n, nil
}

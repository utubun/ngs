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
		read := r.scanner.Bytes()
		text := string(read)
		switch i {
		case 0:
			if !isHeader(text) {
				return 0, fmt.Errorf("invalid format: expected header record, but found %s", text)
			}
			address = int8(len(temp) + len(read))
			temp[1] = byte(address)
			temp = append(temp, read...)
			continue
		case 1:
			if !isDNA(text) {
				return 0, fmt.Errorf("invalid format: expected dna record, but found %s", text)
			}
			address = int8(int(temp[1]) + len(read))
			temp[2] = byte(address)
			temp = append(temp, read...)
			continue
		case 3:
			if !isQualityString(text) {
				return 0, fmt.Errorf("invalid format: expected quality record, but found %s", text)
			}
			temp[0] = byte(1)
			temp = append(temp, read...)
		default:
			continue
		}
	}
	n += copy(b, temp)
	return n, nil
}

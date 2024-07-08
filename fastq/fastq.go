package fastq

import (
	"fmt"
	"io"

	"github.com/utubun/ngs/fastq/internal"
)

type Record struct {
	B []byte
}

func NewRecord(n int) *Record {
	return &Record{B: make([]byte, n)}
}

// Reader reads fastq file
type Reader io.Reader

// NewReader returns new reader
func NewReader(r io.Reader) (Reader, error) {
	return internal.NewReader(r)
}

// DNA gets dna from read
func (r *Record) DNA() (string, error) {
	if r.B[0] == 0 {
		return "", fmt.Errorf("invalid record")
	}
	return string(r.B[r.B[1]:r.B[2]]), nil
}

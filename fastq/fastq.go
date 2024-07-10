package fastq

import (
	"io"
)

type Read interface {
	Len() float64
	Quality() []float64
	Sequence() []rune
}

// Reader reads fastq file
type Reader interface {
	Read() (chan *read, error)
}

// NewReader returns new reader
func NewReader(r io.Reader) Reader {
	reader := newReader(r)
	return reader
}

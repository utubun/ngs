package fastq

type Base struct {
	val  *rune
	qual *int
}

type Read struct {
	pos []*Base
}

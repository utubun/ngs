package internal

type base struct {
	val  *rune
	qual *int
}

type read struct {
	pos []*base
}

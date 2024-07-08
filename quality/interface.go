package quality

type Record interface {
	Length() int
	Validate() bool
	Composition() map[string]float32
}

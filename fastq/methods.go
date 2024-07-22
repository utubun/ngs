package fastq

func newBase(v *rune, q *int) *Base {
	return &Base{
		val:  v,
		qual: q,
	}
}

func newRead() *Read {
	return &Read{}
}

func (r *Read) append(b *Base) {
	r.pos = append(r.pos, b)
}

func (r *Read) Len() float64 {
	return float64(len(r.pos))
}

func (r *Read) Quality() []float64 {
	var res []float64
	for _, val := range r.pos {
		res = append(res, float64(*val.qual))
	}
	return res
}

func (r *Read) Sequence() []rune {
	var res []rune
	for _, val := range r.pos {
		res = append(res, *val.val)
	}
	return res
}

package internal

func newBase(v *rune, q *int) *base {
	return &base{
		val:  v,
		qual: q,
	}
}

func newRead() *read {
	return &read{}
}

func (r *read) append(b *base) {
	r.pos = append(r.pos, b)
}

func (r *read) Len() float64 {
	return float64(len(r.pos))
}

func (r *read) Quality() []float64 {
	var res []float64
	for _, val := range r.pos {
		res = append(res, float64(*val.qual))
	}
	return res
}

func (r *read) Sequence() []rune {
	var res []rune
	for _, val := range r.pos {
		res = append(res, *val.val)
	}
	return res
}

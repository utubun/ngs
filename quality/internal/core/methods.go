package core

func (b *Base) quality() float64 {
	return float64(*b.qvalue)
}

func (s *Seq) quality() []float64 {
	var res []float64
	for _, base := range *s {
		res = append(res, base.quality())
	}
	return res
}

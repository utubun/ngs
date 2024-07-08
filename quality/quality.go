package quality

type Base struct {
	N      int
	QScore []int
}
type QC struct {
	Count int
	A     Base
	C     Base
	T     Base
	G     Base
	N     Base
}

func (q *QC) Check(b []byte) {
	if int(b[0]) == 1 {
		q.Count += 1
	} else {
		return
	}
	dna := b[b[1]:b[2]]
	qual := b[b[2]:]
	if len(dna) != len(qual) {
		return
	}
	for i, v := range dna {
		switch string(v) {
		case "A":
			q.A.N += 1
			q.A.QScore = append(q.A.QScore, int(qual[i]-30))
		case "C":
			q.C.N += 1
			q.C.QScore = append(q.C.QScore, int(qual[i]-30))
		case "T":
			q.T.N += 1
			q.T.QScore = append(q.T.QScore, int(qual[i]-30))
		case "G":
			q.G.N += 1
			q.G.QScore = append(q.G.QScore, int(qual[i]-30))
		default:
			q.N.N += 1
			q.N.QScore = append(q.N.QScore, int(qual[i]-30))

		}
	}
}

package quality

import (
	"bufio"
)

func Scan(s *bufio.Scanner, seq chan string, qual chan string) {
	if ok := s.Scan(); ok {
		str := s.Text()
		switch str[0] {
		case '@':
			s.Scan()
		}
	}
}

func (r Reads) CheckQuality() (FastqQuality, error) {
	var (
		qcheck FastqQuality
		maxLen int
	)

	/*for _, val := range r.Sequence {
		if !isDNA(&val) {
			qcheck.Valid = false
			break
		}
	}*/

	for _, val := range r.QScores {
		/*if !qcheck.Valid {
			break
		}*/
		for _, el := range val {
			if el > 40 || el < 0 {
				qcheck.Valid = false
				break
			}
		}
	}

	for _, val := range r.Sequence {
		i := len(val)

		if i > maxLen {
			maxLen = i
		}

		//qcheck.SeqLength[i]++
	}

	// transpose qualities
	var tqual = make([][]int, maxLen)

	for _, val := range r.QScores {
		for i, q := range val {
			tqual[i] = append(tqual[i], q)
		}
	}

	var qres QStats
	for _, val := range tqual {
		qres.Mean = append(qres.Mean, Mean(val))
		qres.SD = append(qres.SD, SD(val))
	}

	qcheck.Stats = qres

	qcheck.ReadsCount = r.count()

	return qcheck, nil
}

package fastq

import (
	"bufio"
	"compress/gzip"
	"io"
)

func ReadFastq2(r *io.Reader) (Reads, error) {
	var (
		reads Reads
	)

	reader, err := gzip.NewReader(*r)
	if err != nil {
		return reads, err
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	var previous int

	for scanner.Scan() {

		s := scanner.Text()

		switch s[0] {
		case '@':
			previous = 1
			continue
		case '+':
			previous = 2
			continue
			//reads.Sequence = append(reads.Sequence, scanner.Text())
		default:
			switch previous {
			case 1:
				reads.Sequence = append(reads.Sequence, s)
			case 2:
				reads.QScores = append(reads.QScores, convertQualities(s))
			}
		}
	}

	return reads, nil
}

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

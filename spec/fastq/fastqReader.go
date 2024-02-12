package fastq

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
)

func ReadFastq2(b *io.ReadCloser) (Reads, error) {
	var (
		reads Reads
		n     int
	)

	r, err := gzip.NewReader(*b)
	if err != nil {
		return reads, err
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {

		switch n {
		case 0:
			reads.Header = append(reads.Header, scanner.Text())
		case 1:
			reads.Sequence = append(reads.Sequence, scanner.Text())
		case 2:
			n++
			continue
		case 3:
			qScores := convertQualities(scanner.Text())
			reads.QScores = append(reads.QScores, qScores)
		default:
			log.Fatal("Unexpected line")
		}

		n = (n + 1) % 4

	}

	return reads, nil
}

func (r Reads) CheckQuality() (FastqQuality, error) {
	var (
		qcheck FastqQuality
		maxLen int
	)

	for i, val := range r.Sequence {
		if !isDNA(&val) {
			qcheck.Valid = false
			qcheck.Message = append(qcheck.Message, fmt.Sprintf("%s is not DNA", r.Header[i]))
			break
		}
	}

	for i, val := range r.QScores {
		if !qcheck.Valid {
			break
		}
		for _, el := range val {
			if el > 40 || el < 0 {
				qcheck.Valid = false
				qcheck.Message = append(qcheck.Message, fmt.Sprintf("%s invalid quality score %d", r.Header[i], el))
			}
			break
		}
	}

	for _, val := range r.Sequence {
		i := len(val)

		if i > maxLen {
			maxLen = i
		}

		qcheck.SeqLength[i]++
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

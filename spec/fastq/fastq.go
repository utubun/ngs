package fastq

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"log"
	"os"
)

// ReadZip function
// Reads zipped fastq file
func ReadFastq() FastqQuality {
	var (
		reads  Reads
		qcheck FastqQuality
		n      int
	)

	qcheck.Valid = true
	qcheck.SeqLength = make(map[int]int)

	buf, err := os.Open("../testdata/SRR24211928_R1.fastq.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer buf.Close()

	f, err := gzip.NewReader(buf)
	if err != nil {
		qcheck.Valid = false
		qcheck.Message = append(qcheck.Message, fmt.Sprintf("Not valid application/gzip file"))
		return qcheck
	}
	defer f.Close()

	fmt.Printf("Contents of %s\n", f.Name)

	scanner := bufio.NewScanner(f)
	//scanner.Split(bufio.ScanLines)

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

	for i, val := range reads.Sequence {
		if !isDNA(&val) {
			qcheck.Valid = false
			qcheck.Message = append(qcheck.Message, fmt.Sprintf("%s is not DNA", reads.Header[i]))
			break
		}
	}

	for i, val := range reads.QScores {
		if !qcheck.Valid {
			break
		}
		for _, el := range val {
			if el > 40 || el < 0 {
				qcheck.Valid = false
				qcheck.Message = append(qcheck.Message, fmt.Sprintf("%s invalid quality score %d", reads.Header[i], el))
			}
			break
		}
	}

	var maxLen int

	for _, val := range reads.Sequence {
		i := len(val)

		if i > maxLen {
			maxLen = i
		}

		qcheck.SeqLength[i]++
	}

	// transpose qualities
	var tqual = make([][]int, maxLen)

	for _, val := range reads.QScores {
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

	qcheck.ReadsCount = reads.count()

	return qcheck

}

package quality

/*
import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func ReadLocalFile(p string) (io.Reader, error) {
	buf, err := os.Open(p)

	if err != nil {
		return nil, err
	}

	return buf, err
}

func ReadGz(buf io.Reader) (*bufio.Scanner, error) {

	r, err := gzip.NewReader(buf)
	defer r.Close()
	if err != nil {
		fmt.Printf("Error reading archive: %+v", err)
		return nil, err
	}

	s := bufio.NewScanner(r)
	return s, nil
}

func Read(r io.Reader) (Reads, error) {
	var (
		reads Reads
	)

	reader, err := gzip.NewReader(r)
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
*/

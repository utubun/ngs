package fastq

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

func Counter2(s *bufio.Scanner) int {
	var (
		n int
		c int
	)

	for s.Scan() {

		/*switch n {
		case 0:
			//reads.Header = append(reads.Header, scanner.Text())
			fmt.Println("Foun")
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
		}*/

		n = (n + 1) % 4
		c++

	}

	return c
}

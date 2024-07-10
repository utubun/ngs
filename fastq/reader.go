package fastq

import (
	"bufio"
	"fmt"
	"io"
	"sync"
)

// alphabet ACGNT 65 67 71 78 84

type reader struct {
	alphabet [5]rune
	quality  [65]int
	scanner  *bufio.Scanner
}

func newReader(r io.Reader) *reader {
	alphabet := [5]rune{65, 67, 71, 78, 84}
	var quality [65]int
	for i := range quality {
		quality[i] = i
	}
	scanner := bufio.NewScanner(r)
	return &reader{alphabet: alphabet, quality: quality, scanner: scanner}
}

func (r *reader) getRecord(qbase int, ch chan *read, wg *sync.WaitGroup) error {
	// define dna sequence
	var seq string
	// define quality sequence
	var qual string
	for i := 0; i < 4; i++ {
		ok := r.scanner.Scan()
		if !ok {
			wg.Done()
			return io.EOF
		}
		text := r.scanner.Text()
		switch i {
		case 0:
			if !isHeader(text) {
				wg.Done()
				return fmt.Errorf("invalid format: expected header record, but found %s", text)
			}
			continue
		case 1:
			if !isDNA(text) {
				wg.Done()
				return fmt.Errorf("invalid format: expected dna record, but found %s", text)
			}
			seq = text
			continue
		case 3:
			if !isQualityString(text) {
				wg.Done()
				return fmt.Errorf("invalid format: expected quality record, but found %s", text)
			}
			qual = text
		default:
			continue
		}
	}
	// prepare and send the read
	go func(seq, qual string, wg *sync.WaitGroup) {
		defer wg.Done()
		// define result
		res := newRead()
		for i, val := range seq {
			base := &base{}
			switch val {
			case 65:
				base.val = &r.alphabet[0]
			case 67:
				base.val = &r.alphabet[1]
			case 71:
				base.val = &r.alphabet[2]
			case 84:
				base.val = &r.alphabet[4]
			default:
				base.val = &r.alphabet[4]
			}
			base.qual = &r.quality[int(qual[i])-qbase]
			res.append(base)
		}
		ch <- res
	}(seq, qual, wg)
	return nil
}

func (r *reader) Read() (chan *read, error) {
	//
	var wg sync.WaitGroup
	// define output chanel
	ch := make(chan *read)
	// iterate over input and write data to the channel
	for {
		wg.Add(1)
		err := r.getRecord(33, ch, &wg)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch, nil
}

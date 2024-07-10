package core

import (
	"math"
	"sync"
)

const (
	BASE = 33
)

/*
	func ProcessSequence(b []byte, seqID int, input *QC) (*QC, error) {
		// check the record is valid fastq record
		if int(b[0]) != 1 {
			return nil, fmt.Errorf("invalid fastq record")
		}
		// parse dna and quality score records
		dna := b[b[1]:b[2]]
		qual := b[b[2]:]
		// iterate over dna byte array and populate QC instance with data
		for i, v := range dna {
			nt := string(v)           // get nucleotide
			qv := int(qual[i]) - BASE // get quality score
			// populate array of q-values with
			input.Qual[qv] = qv
			// check that QC.Seq has enough capacity and extend if necessary
			if len(input.Seq) <= seqID {
				input.Seq = input.Seq[:seqID+1]
			}
			// append base containing pointer to q-value to the sequence
			input.Seq[seqID] = append(input.Seq[seqID], &Base{Quality: &input.Qual[qv]})
			// update base count
			if entry, ok := input.Bases[nt]; ok {
				if len(entry) <= seqID {
					entry = append(entry, make([]int, seqID-len(entry)+1)...)
				}
				entry[seqID] += 1
				input.Bases[nt] = entry
			} else {
				entry := make([]int, seqID+1)
				entry[0] = 1
				input.Bases[nt] = entry
			}
		}
		return input, nil
	}

	func QualityPerPosition(input *QC) []Point {
		// initialize point array as result
		res := make([]Point, 0)
		// initialize point array maximal length with 0 value
		var max int
		// find the the seq with max length, and assign to max
		for _, val := range input.Seq {
			l := len(val)
			if max < l {
				max = l
			}
		}
		// collect the data for each position [[pos 1 seq 1, pos 2 seq 2...], [pos 2 seq 1, pos2 seq 2, ...], ...]
		y := make([][]int, max)
		for _, val := range input.Seq {
			for j, base := range val {
				y[j] = append(y[j], *base.Quality)
			}
		}
		// prepare final results
		for i, val := range y {
			// create point instance
			p := &Point{
				X:     float64(i), // index of outer array is the position
				Y:     Mean(val),  // calculate mean quality score
				Label: "Quality Score",
			}
			res = append(res, *p)
		}

		return res
	}

	func SeqLenDistribution(input *QC) []float64 {
		// initialize array as a result
		res := make([]float64, len(input.Seq))
		for i, val := range input.Seq {
			res[i] = float64(len(val))
		}
		return res
	}

	func PerSeqQuality(input *QC) []float64 {
		// initialize point array as result
		res := make([]float64, len(input.Seq))
		// iterate over seq and find mean score per sequence
		for i, val := range input.Seq {
			y := make([]int, len(val))
			for j, base := range val {
				y[j] = *base.Quality
			}
			res[i] = Mean(y)
		}
		return res
	}

	func PerSeqGC(input *QC) []float64 {
		// initialize point array as result
		res := make([]float64, len(input.Seq))
		// collect the data for each position [[pos 1 seq 1, pos 2 seq 2...], [pos 2 seq 1, pos2 seq 2, ...], ...]
		total := SeqLenDistribution(input)
		for base, val := range input.Bases {
			for j, n := range val {
				if base == "C" || base == "G" {
					res[j] += float64(n) / total[j] * 100.0
				}
			}
		}
		return res
	}
*/
func sum[T int | int32 | int64 | float32 | float64](x []T) float64 {
	var sum float64
	for _, val := range x {
		sum += float64(val)
	}

	return float64(sum)
}

func mean[T int | int32 | int64 | float32 | float64](x []T) float64 {
	s := sum(x)
	l := float64(len(x))
	return s / l
}

func dev[T int | int32 | int64 | float32 | float64](x []T) float64 {
	var deviation float64
	m := mean(x)
	for _, val := range x {
		deviation += math.Pow(float64(val)-m, float64(2))
	}
	return deviation
}

func sd[T int | int32 | int64 | float32 | float64](x []T) float64 {
	deviation := math.Sqrt(dev(x) / float64(len(x)))
	return deviation
}

func max[T int | int32 | int64 | float32 | float64](x []T) float64 {
	var res float64
	for _, val := range x {
		res = math.Max(res, float64(val))
	}
	return res
}

func min[T int | int32 | int64 | float32 | float64](x []T) float64 {
	var res float64
	for _, val := range x[1:] {
		res = math.Min(res, float64(val))
	}
	return res
}

func Summary[T Number](x []T) *Statistics[T] {
	// waitgroup and lock
	var wg sync.WaitGroup
	var lock sync.Mutex
	// define result
	res := &Statistics[T]{Data: x}
	// check if x is empty
	if len(x) == 0 {
		return &Statistics[T]{}
	}
	// calculate length of data
	wg.Add(1)
	go func() {
		lock.Lock()
		defer lock.Unlock()
		defer wg.Done()
		res.N = float64(len(x))
	}()
	// calculate min value
	wg.Add(1)
	go func() {
		lock.Lock()
		defer lock.Unlock()
		defer wg.Done()
		res.Min = min(x)
	}()
	// calculate max value
	wg.Add(1)
	go func() {
		lock.Lock()
		defer lock.Unlock()
		defer wg.Done()
		res.Max = max(x)
	}()
	// calculate mean value
	wg.Add(1)
	go func() {
		lock.Lock()
		defer lock.Unlock()
		defer wg.Done()
		res.Mean = mean(x)
	}()
	// calculate sd value
	wg.Add(1)
	go func() {
		lock.Lock()
		defer lock.Unlock()
		defer wg.Done()
		res.SD = sd(x)
	}()
	wg.Wait()
	return res
}

func Count(s string) map[rune]int64 {
	// define result
	res := make(map[rune]int64)
	for _, char := range s {
		res[char] += 1
	}
	return res
}

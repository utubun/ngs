package fastq

import (
	"math"
	"regexp"
)

type Header string

type Sequence string

type QScores []int

type Reads struct {
	Header   []string
	Sequence []string
	QScores  [][]int
}

type QStats struct {
	Count  int
	Mean   []float64
	Median float64
	SD     []float64
	Min    []float64
	Max    []float64
}

type FastqQuality struct {
	Valid      bool        `json:"isValid" default:"true"`
	Filetype   string      `json:"fileType" default:"octet-stream"`
	ReadsCount int         `json:"count" default:"0"`
	SeqLength  map[int]int `json:"seqLength"`
	Stats      QStats      `json:"stats"`
	Message    []string    `json:"message"`
}

func Sum(x []int) int {
	var sum int
	for _, val := range x {
		sum += val
	}

	return sum
}

func Mean(x []int) float64 {
	sum := Sum(x)
	avr := float64(sum / len(x))
	return avr
}

func Deviation(x []int) float64 {
	var deviation float64
	mean := Mean(x)
	for _, val := range x {
		deviation += math.Pow(float64(val)-mean, float64(2))
	}
	return deviation
}

func SD(x []int) float64 {
	dev := math.Sqrt(Deviation(x) / float64(len(x)))
	return dev
}

func Max(x []int) float64 {
	var max float64
	for _, val := range x {
		max = math.Max(max, float64(val))
	}
	return max
}

func Min(x []int) float64 {
	var min float64
	for _, val := range x {
		min = math.Min(min, float64(val))
	}
	return min
}

func (r *Reads) count() int {
	return len(r.Header)
}

var dnaAlphabet = regexp.MustCompile(`^[GATCN]+$`).MatchString

func isDNA(s *string) bool {
	return dnaAlphabet(*s)
}

func (r *Reads) QualityByPositionStats() QStats {

	return QStats{}
}

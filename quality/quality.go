package quality

import (
	"fmt"
	"io"
	"log"
	"math"
	"slices"

	"github.com/utubun/ngs/fastq"
	"gonum.org/v1/gonum/stat"
)

func QC(r io.Reader) (*Report, error) {
	fq := fastq.NewReader(r)
	reads, err := fq.Read()
	if err != nil {
		return nil, err
	}
	report := NewReport()
	for read := range reads {
		report.Update(*read)
	}

	return report, nil
}

type Report struct {
	Encoding        int          `json:"encoding"`
	N               int          `json:"count"`
	Length          Statistics   `json:"length"`
	Quality         Statistics   `json:"quality"`
	QualPerPosition []Statistics `json:"quality_per_position"`
	GC              Statistics   `json:"gc"`
	GCPerSeq        []float64    `json:"gc_per_seq`
}

func NewReport() *Report {
	return &Report{
		QualPerPosition: make([]Statistics, 1),
	}
}

func (r *Report) Update(input fastq.Read) {
	r.N += 1
	r.Length.Update(input.Len())

	// update quality score statistics
	qual := input.Quality()
	r.Quality.Update(stat.Mean(qual, nil))
	// extend the array of quality per position if needed
	if qplen := len(r.QualPerPosition); qplen < len(qual) {
		r.QualPerPosition = append(r.QualPerPosition, make([]Statistics, len(qual)-qplen)...)
	}
	for i, q := range input.Quality() {
		r.QualPerPosition[i].Update(q)
	}
	gc := GC(input.Sequence())
	r.GC.Update(gc)

	r.GCPerSeq = append(r.GCPerSeq, gc)

}

type Statistics struct {
	N    float64   `json:"n"`
	Min  float64   `json:"min"`
	Max  float64   `json:"max"`
	Mean float64   `json:"mean"`
	Var  float64   `json:"variance"`
	SD   float64   `json:"stdev"`
	Q    []float64 `json:"quartile"`
	Out  []float64 `json:"outliers"`
	m2   float64
}

func (s *Statistics) Update(x float64) {
	// update statistics
	// see https://en.wikipedia.org/wiki/Algorithms_for_calculating_variance
	if s.N == 0 {
		s.Min = x
		s.Max = x
	}
	s.N += 1
	// Calculate mean
	delta := x - s.Mean
	s.Mean += delta / s.N
	// Prepare SD calculation
	s.m2 = delta * (x - s.Mean)
	// Update min, sd, variance
	if s.N > 1 {
		s.Min = math.Min(x, s.Min)
		s.Max = math.Max(x, s.Max)
		s.Var = s.m2 / (s.N - 1)
		s.SD = math.Sqrt(s.Var)
		// calculate quartiles
		s.Q = []float64{s.Min, s.Mean - 0.675*s.SD, s.Mean, s.Mean + 0.675*s.SD, s.Max}
		out := s.Out
		s.Out = []float64{}
		for _, val := range out {
			if val < s.Q[1] && val > s.Q[3] {
				s.Out = append(s.Out, val)
			}
		}
	}
}

func GC(seq []rune) float64 {
	count := len(seq)
	if count == 0 {
		return 0
	}
	for _, base := range seq {
		if base != 'G' && base != 'C' {
			count -= 1
		}
	}
	return float64(count) / float64(len(seq))
}

type Point struct {
	X     float64                `json:"x"`
	Y     float64                `json:"y"`
	Label string                 `json:"label"`
	Props map[string]interface{} `json:"props" default:"nil"`
}

type Histogram []Point

func NewHistogram(x []float64, bw float64) []Point {
	slices.Sort(x)
	l := len(x)
	var h float64

	if bw == 0 {
		iqr := stat.Quantile(.75, stat.Empirical, x, nil) - stat.Quantile(.25, stat.Empirical, x, nil)
		log.Printf("IQR: %.2f\n", iqr)
		h = 2 * iqr / math.Pow(float64(l), 1.0/3.0)
		log.Printf("Bandwidth: %.02f\n", h)
	} else {
		h = float64(x[len(x)-1]-x[0]) / bw
	}

	pnt := Point{X: h + x[0]}
	var hist []Point
	for _, v := range x {
		if v < pnt.X {
			pnt.Y += 1.0
			pnt.Label = fmt.Sprintf("X: %.02f, Y: %.02f", pnt.X, pnt.Y)
		} else {
			hist = append(hist, pnt)
			pnt.X += h
			pnt.Y = 1.0
			pnt.Label = fmt.Sprintf("x: %.02f, y: %.02f", pnt.X, pnt.Y)
		}
	}
	return hist
}

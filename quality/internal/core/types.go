package core

type Report struct {
	Encoding              int        `json:"encoding"`
	NumberOfReads         int        `json:"nreads"`
	SeqLengthDistribution Statistics `json:"length"`
	QualityPerPosition    Statistics `json:"qposition"`
	QualityPerSequence    Statistics `json:"qseq"`
	GCPerPosition         Statistics `json:"gcposition"`
	GCPerSequence         Statistics `json:"gcseq"`
}

type Base struct {
	qvalue *int
}

type Seq []*Base

type Sequence []*Base

type QC struct {
	Bases map[string][]int
	Seq   []Sequence
	Qual  [100]int
}

type Number interface{ int | int64 | float64 }

type Point[T Number] struct {
	X     T                      `json:"x default:"null""`
	Y     T                      `json:"y" default:"null"`
	Label string                 `json:"label" default:"null"`
	Props map[string]interface{} `json:"props" default:"null"`
}

type Statistics struct {
	N        float64   `json:"n"`
	Min      float64   `json:"min"`
	Max      float64   `json:"max"`
	Mean     float64   `json:"mean"`
	Var      float64   `json:"variance"`
	SD       float64   `json:"stdev"`
	Outliers []float64 `json:"outliers"`
	m2       float64
}

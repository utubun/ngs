package core

type Report struct {
	Encoding              int                 `json:"encoding"`
	NumberOfReads         int                 `json:"nreads"`
	SeqLengthDistribution Statistics[int64]   `json:"length"`
	QualityPerPosition    Statistics[float64] `json:"qposition"`
	QualityPerSequence    Statistics[float64] `json:"qseq"`
	GCPerPosition         Statistics[float64] `json:"gcposition"`
	GCPerSequence         Statistics[float64] `json:"gcseq"`
}

type Base struct {
	qvalue *int
	child  *Base
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

type Statistics[T Number] struct {
	N    float64 `json:"n"`
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
	Mean float64 `json:"mean"`
	SD   float64 `json:"std"`
	Data []T     `json:"val"`
}

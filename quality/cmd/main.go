package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/utubun/ngs/fastq"
	"github.com/utubun/ngs/quality/internal/util"
)

func main() {
	f, err := os.Open("../quality/internal/assets/short.fastq")

	if err != nil {
		log.Printf("Error opening th efile: %s", err)
	}
	defer f.Close()
	r, err := fastq.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	b := make([]byte, 1024)
	q := NeqQC()
	for {
		_, err := r.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		q.Check(b)
	}
	gc := q.GC()
	fmt.Println(gc)
	qpos := q.QualityPerPosition()
	fmt.Println(qpos)
}

type base struct {
	Val   string
	Count int
}

type quality struct {
	Val   int
	Count int
}

type position struct {
	Base    base
	Quality quality
}

type Point struct {
	X     int64
	Y     int64
	Label string
	Props map[string]interface{}
}

type QC struct {
	Count    int
	Len      []int
	Base     map[string]base
	Qual     map[int]quality
	Position map[int][]position
}

func NeqQC() *QC {
	return &QC{
		Position: make(map[int][]position),
		Base:     make(map[string]base),
		Qual:     make(map[int]quality),
	}
}

func (q *QC) Check(b []byte) {
	if int(b[0]) == 1 {
		q.Count += 1
	} else {
		return
	}
	dna := b[b[1]:b[2]]
	qual := b[b[2]:]
	q.Len = append(q.Len, len(dna))
	for i, v := range dna {
		nt := string(v)
		qv := int(qual[i])
		if entry, ok := q.Base[nt]; ok {
			entry.Count += 1
			q.Base[nt] = entry
		} else {
			entry = base{nt, 1}
			q.Base[nt] = entry
		}
		if entry, ok := q.Qual[qv]; ok {
			entry.Count += 1
			q.Qual[qv] = entry
		} else {
			entry = quality{qv, 1}
			q.Qual[qv] = entry
		}

		q.Position[i] = append(q.Position[i], position{q.Base[nt], q.Qual[qv]})
	}
}

func (q *QC) GC() map[string]float64 {
	m := make(map[string]float64)
	var total float64
	for _, val := range q.Base {
		total += float64(val.Count)
	}
	for key, val := range q.Base {
		m[key] = float64(val.Count) / total * 100
	}
	return m
}

func (q *QC) QualityPerPosition() *[]Point {
	var res []Point
	for key, val := range q.Position {
		var y []int
		for _, p := range val {
			y = append(y, p.Quality.Val)
		}
		p := &Point{
			X: int64(key),
			Y: int64(util.Mean(y)),
		}
		res = append(res, *p)
	}
	return &res
}

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/utubun/ngs/quality/internal/core"
)

const (
	_ = 16 << iota
	BASE32
	BASE64
)

func main() {
	/* f, err := os.Open("../quality/internal/assets/tiny.fastq")

	if err != nil {
		log.Printf("Error opening th efile: %s", err)
	}
	defer f.Close()
	r, err := fastq.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	b := make([]byte, 1024)
	q := quality.NewQC()

	var id int
	for {
		_, err := r.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		q, err = core.ProcessSequence(b, id, q)
		if err != nil {
			log.Fatal(err)
		}
		id += 1
	}

	//fmt.Printf("%#v\n", q)

	pq := core.QualityPerPosition(q)
	//fmt.Printf("%+v\n", pq)
	/*gc := q.GC()
	fmt.Println(gc)

	js, _ := json.Marshal(pq)
	os.WriteFile("qualpp.json", js, os.ModePerm)

	pp := core.PerSeqQuality(q)
	js, _ = json.Marshal(pp)
	os.WriteFile("qualps.json", js, os.ModePerm)

	//ld := core.SeqLenDistribution(q)
	//fmt.Printf("Seq Length Distribution: %v", ld)

	gc := core.PerSeqGC(q)
	//fmt.Printf("GC content per sequence: %v\n", gc)
	js, _ = json.Marshal(gc)
	os.WriteFile("gc.json", js, os.ModePerm)

	fmt.Printf("BASE32 is: %d. BASE64 is %d", BASE32, BASE64) */
	report := core.Report{}
	dat := make([]core.Seq, 109)

	for i := 0; i < 109; i++ {
		l := rand.Intn(109)
		for j := 0; j < l; j++ {
			dat[i] = append(dat[i], &core.Base{})
		}
	}
	report.Make(dat)

	s := "ACCGTCGTTTCGAAAAAAAAANA"
	count := core.Count(s)
	for key, val := range count {
		fmt.Println(string(key), ": ", val)
	}
	time.Sleep(5 * time.Second)
	fmt.Printf("Summary:\n%+v\n", report)
}

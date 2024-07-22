package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/utubun/ngs/fastq"
	"github.com/utubun/ngs/quality"
)

const (
	_ = 16 << iota
	BASE32
	BASE64
)

func main() {
	//var wg sync.WaitGroup
	//var lock sync.Mutex

	start := time.Now()

	f, err := os.Open("../quality/internal/assets/big.fastq")
	if err != nil {
		log.Printf("Error opening th efile: %s", err)
	}
	defer f.Close()

	fmt.Printf("Read the file in %s\n", time.Since(start))

	start = time.Now()

	report := quality.NewReport()

	r := fastq.NewReader(f)
	ch, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created the reader, and sent all the data to the chanel in %s\n", time.Since(start))

	start = time.Now()

	for input := range ch {
		report.Update(*input)
	}

	fmt.Printf("Updated the report in %s\n", time.Since(start))

	//start = time.Now()

	/*js, _ := json.Marshal(report)

	fmt.Printf("Serialized and printed the data in %s\n", time.Since(start))
	fmt.Printf("Length of the Quality per Position: %d\n", len(report.QualPerPosition))
	*/

	/*hist := quality.NewHistogram(report.GCPerSeq, 0)

	fmt.Println("Histogram:")
	for _, val := range hist {
		fmt.Printf("x: %.02f, y: %.02f\n", val.X, val.Y)
	}
	js, _ = json.Marshal(hist)
	os.WriteFile("gc.json", js, os.ModePerm)*/
	fmt.Printf("Quality:\n%+v\n", report.Quality)
}

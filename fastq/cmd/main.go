package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/utubun/ngs/fastq"
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
	for {
		_, err := r.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		rec := &fastq.Record{B: b}
		dna, _ := rec.DNA()
		fmt.Printf("DNA string:\n%s\n", dna)
	}
}

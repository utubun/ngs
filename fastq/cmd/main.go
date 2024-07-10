package main

import (
	"fmt"
	"log"
	"os"

	"github.com/utubun/ngs/fastq"
)

func main() {
	f, err := os.Open("../quality/internal/assets/big.fastq")

	if err != nil {
		log.Printf("Error opening th efile: %s", err)
	}
	defer f.Close()
	r := fastq.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	ch, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}
	for read := range ch {
		fmt.Printf("Read:\n%s\n", string(read.Sequence()))
	}

}

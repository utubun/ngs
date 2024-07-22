package core

import (
	"sync"

	"github.com/utubun/ngs/fastq"
)

func Make(reads chan *fastq.Read) *Report {
	// wait group
	var wg sync.WaitGroup
	// define result
	report := &Report{}
	// define chanel for length distribution
	length := make(chan int64)
	// define chanel for quality distribution
	var quality []float64
	//var stream stat.StreamStats
	//github.com/aclements/go-moremath/stats
	// iterate over reader
	for read := range reads {
		// record number of reads
		report.NumberOfReads += 1

		// record length distribution
		wg.Add(1)
		go func(read *fastq.Read) {
			defer wg.Done()
			length <- int64(read.Len())
		}(read)

		quality = append(quality, mean(read.Quality()))

	}

	go func() {
		wg.Wait()
		close(length)
	}()

	var lenDistr []int64
	for l := range length {
		lenDistr = append(lenDistr, l)
	}
	report.SeqLengthDistribution = *Summary(lenDistr)
	report.QualityPerSequence = *Summary(quality)

	return report
}

func (r *Report) seqLenDistribution(s []*Seq) {
	// define wait group
	var wg sync.WaitGroup
	// define mutex
	var lock sync.Mutex
	// run go routine
	wg.Add(1)
	go func(sequence []*Seq) {
		lock.Lock()
		defer lock.Unlock()
		defer wg.Done()
		// prepare data
		var dat []int64
		for _, seq := range sequence {
			dat = append(dat, int64(len(*seq)))
		}
		r.SeqLengthDistribution = *Summary(dat)
	}(s)
	// wait on background
	go func() {
		wg.Wait()
	}()
}

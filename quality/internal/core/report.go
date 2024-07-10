package core

import (
	"fmt"
	"sync"
)

func (r *Report) Make(s []Seq) {
	// define wait group
	var wg sync.WaitGroup
	// define mutex
	var lock sync.Mutex
	// run go routine
	wg.Add(1)
	go func(sequence []Seq) {
		lock.Lock()
		defer lock.Unlock()
		defer wg.Done()
		// prepare data
		var dat []int64
		for _, seq := range sequence {
			dat = append(dat, int64(len(seq)))
		}
		fmt.Println(dat)
		r.SeqLengthDistribution = *Summary(dat)
	}(s)
	// wait on background
	go func() {
		wg.Wait()
	}()
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

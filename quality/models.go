package quality

import (
	"math"
	"regexp"
	"strconv"
	"sync"
	"sync/atomic"
)

var (
	// "@M07197:20:000000000-K7JRN:1:1101:16242:1028 1:N:0:8"
	headerRe = regexp.MustCompile(`^\@(?P<Instrument>[A-Za-z0-9]+):(?P<Run>\d+):(?P<Flowcell>[A-Za-z0-9\-]+):(?P<Lane>\d?):(?P<Tile>\d+):(?P<X>\d+):(?P<Y>\d+)\s(?P<Read>\d+):(?P<Filtered>[YN]):(?P<Control>\d?):(?P<Sample>\d?)$`)
)

type RawHeader string

type Header struct {
	Instrument string `json:"instrument" default:"unknown"`
	Run        int    `json:"run"`
	Flowcell   string `json:"flowcell" default:"unknown"`
	Lane       int    `json:"lane"`
	Tile       int    `json:"tile"`
	Coord      []int  `json:"coordinates"`
	Read       int    `json:"read"`
	Control    int    `json:"control"`
	Sample     int    `json:"sample"`
	Filtered   bool   `json:"filtered"`
}

func ParseHeader(h string) Header {
	var header Header

	match := headerRe.FindStringSubmatch(h)

	header.Instrument = match[headerRe.SubexpIndex("Instrument")]
	header.Run, _ = strconv.Atoi(match[headerRe.SubexpIndex("Run")])
	header.Flowcell = match[headerRe.SubexpIndex("Flowcell")]
	header.Lane, _ = strconv.Atoi(match[headerRe.SubexpIndex("Lane")])
	header.Tile, _ = strconv.Atoi(match[headerRe.SubexpIndex("Tile")])
	x, _ := strconv.Atoi(match[headerRe.SubexpIndex("X")])
	y, _ := strconv.Atoi(match[headerRe.SubexpIndex("Y")])
	header.Coord = []int{x, y}
	header.Read, _ = strconv.Atoi(match[headerRe.SubexpIndex("Read")])
	header.Control, _ = strconv.Atoi(match[headerRe.SubexpIndex("Control")])
	header.Sample, _ = strconv.Atoi(match[headerRe.SubexpIndex("Sample")])
	header.Filtered = match[headerRe.SubexpIndex("Filtered")] == "Y"

	return header
}

type DNA string

func BaseFrequency(dna string) map[rune]int {
	frequency := make(map[rune]int)

	for _, base := range dna {
		frequency[base]++
	}

	return frequency
}

func (r *Reads) BaseFrequency() map[rune]int {
	frequency := make(map[rune]int)

	for _, dna := range r.Sequence {
		freq := BaseFrequency(dna)
		for key, value := range freq {
			frequency[key] += value
		}
	}
	return frequency
}

type Counter struct {
	m sync.Map
}

func (c *Counter) Add(key string, value int64) int64 {
	count, ok := c.m.LoadOrStore(key, &value)
	if ok {
		return atomic.AddInt64(count.(*int64), value)
	}
	return *count.(*int64)
}

func (c *Counter) Load(key string) int64 {
	count, ok := c.m.Load(key)
	if ok {
		return atomic.LoadInt64(count.(*int64))
	}

	return *count.(*int64)
}

func (s *DNA) Validate() bool {

	var alphabet map[string]bool
	alphabet = map[string]bool{
		"A": true,
		"C": true,
		"T": true,
		"G": true,
		"N": true,
	}

	for _, val := range *s {
		if !alphabet[string(val)] {
			return false
		}
	}

	return true
}

func (s *DNA) Composition(c *Counter) {

	var wg sync.WaitGroup

	for _, val := range *s {
		wg.Add(1)

		go func() {
			defer wg.Done()
			go c.Add(string(val), 1)
		}()

		wg.Wait()
	}
}

type Quality string

type QScores []int

type Reads struct {
	Header   []Header
	Sequence []string
	QScores  [][]int
}

type QStats struct {
	Count  int
	Mean   []float64
	Median float64
	SD     []float64
	Min    []float64
	Max    []float64
}

type FastqQuality struct {
	Valid      bool        `json:"isValid" default:"true"`
	Filetype   string      `json:"fileType" default:"octet-stream"`
	ReadsCount int         `json:"count" default:"0"`
	SeqLength  map[int]int `json:"seqLength"`
	Stats      QStats      `json:"stats"`
	Message    []string    `json:"message"`
}

/* func GC(t map[rune]int) float32 {
	var all float32
	gc := float32(t[71]) + float32(t[67])
	for _, val := range t {
		all += float32(val)
	}

	return gc / all * 100
} */

func Sum(x []int) int {
	var sum int
	for _, val := range x {
		sum += val
	}

	return sum
}

func Mean(x []int) float64 {
	sum := Sum(x)
	avr := float64(sum / len(x))
	return avr
}

func Deviation(x []int) float64 {
	var deviation float64
	mean := Mean(x)
	for _, val := range x {
		deviation += math.Pow(float64(val)-mean, float64(2))
	}
	return deviation
}

func SD(x []int) float64 {
	dev := math.Sqrt(Deviation(x) / float64(len(x)))
	return dev
}

func Max(x []int) float64 {
	var max float64
	for _, val := range x {
		max = math.Max(max, float64(val))
	}
	return max
}

func Min(x []int) float64 {
	var min float64
	for _, val := range x {
		min = math.Min(min, float64(val))
	}
	return min
}

func (r *Reads) count() int {
	return len(r.Header)
}

var dnaAlphabet = regexp.MustCompile(`^[GATCN]+$`).MatchString

func isDNA(s *string) bool {
	return dnaAlphabet(*s)
}

func (r *Reads) QualityByPositionStats() QStats {

	return QStats{}
}

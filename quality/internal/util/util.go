package util

import (
	"math"
	"regexp"
)

func isDNA(x string) bool {
	return regexp.MustCompile(`^[ACTGNactgn]+$`).MatchString(x)
}

func isHeader(x string) bool {
	return regexp.MustCompile(`^@[A-Za-z0-9:=._ ]+$`).MatchString(x)
}

func isQualityString(x string) bool {
	return regexp.MustCompile(`^[@#;:+=<>'.)(?!A-Za-z0-9]|-+$`).MatchString(x)
}

func isUtilString(x string) bool {
	return regexp.MustCompile(`^\+.*$`).MatchString(x)
}

func IdentifyReadLine(x string) string {
	if isHeader(x) {
		return "header"
	}
	if isDNA(x) {
		return "dna"
	}
	if isQualityString(x) {
		return "quality"
	}
	if isUtilString(x) {
		return "util"
	}
	return "unknown"
}

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

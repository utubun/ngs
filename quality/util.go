package quality

import (
	"regexp"
)

var (
	pheader = regexp.MustCompile(`^@.+$`)
	pseq    = regexp.MustCompile(`^[ACTG]+$`)
	qheader = regexp.MustCompile(`^\+.*$`)
)

func isHeader(s string) bool {
	return pheader.MatchString(s)
}

func isQHeader(s string) bool {
	return qheader.MatchString(s)
}

func convertQualities(s string) []int {
	var res []int
	for _, val := range s {
		res = append(res, int(val)-33)
	}

	return res
}

func isSeq(s string) bool {
	return pseq.MatchString(s)
}

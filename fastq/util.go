package fastq

import (
	"regexp"
)

func isHeader(x string) bool {
	return regexp.MustCompile(`^@[A-Za-z0-9:=._ ]+$`).MatchString(x)
}

func isDNA(x string) bool {
	return regexp.MustCompile(`^[ACTGNactgn]+$`).MatchString(x)
}

func isUtilString(x string) bool {
	return regexp.MustCompile(`^\+[A-Za-z0-9:_ .=]+$`).MatchString(x)
}

func isQualityString(x string) bool {
	return regexp.MustCompile(`^[@#;:+=<>'.)(?!A-Za-z0-9]+|-+$`).MatchString(x)
}

func IdentifyReadLine(x string) string {
	if isHeader(x) {
		return "header"
	}
	if isDNA(x) {
		return "dna"
	}
	if isUtilString(x) {
		return "util"
	}
	if isQualityString(x) {
		return "quality"
	}
	return "unknown"
}

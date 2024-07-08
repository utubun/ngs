package main

import "fmt"

const (
	seq1 = `ATGGAGGACCGATCATCATCATCAAC`
	seq2 = `ATGGAGGACCGATCATCATCATCATC`
	mix  = `A@T@G@ADTDAFAFGDCHAFAHAHGFCFCGAICJAITITIAGGGAIAJAJAJTJGGTHGHA@CBTHTETIGGCHCGAIAJTJTJTJCGAEGFAGAGTHGIC>TFTGAETGTHGICDAGADABC?A@T@G@ADTDAFAFGDCHAFAHAHGFCFCGAICJAITITIAGGGAIAJAJAJTJGGTHGHA@CBTHTETIGGCHCGAIAJTJTJTJCGAEGFAGAGTHGIC>TFTGAETGTHGICDAGADABC?A@T@G@ADTDAFAFGDCHAFAHAHGFCFCGAICJAITITIAGGGAIAJAJAJTJGGTHGHA@CBTHTETIGGCHCGAIAJTJTJTJCGAEGFAGAGTHGIC>TFTGAETGTHGICDAGADABC?C(G@ABAGAITETHAGTGTIAGGCAHAHAHGGTATCGETDTDGDGDT;C>TCTEGECCT:CATCTDGDACACC:G<G=C=ABT<C9A:G:C@ACGDTCGDC:T@ACT>T@C3A@A@T@TCGCT@TCGBT<TBTCA3A:TCTDTDCCT3T>C>TCA@A4T:ACA#T#T#G#T#G#C#`
	//seq = `AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAATAATTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTGTGC`
)

var (
	enc = map[string]int{
		"C": 0,
		"G": 1,
		"A": 2,
		"T": 3,
		"N": 4,
		"@": 5,
		"D": 6,
		"F": 7,
		"H": 8,
		"I": 9,
		"J": 10,
		"B": 11,
		"E": 12,
		">": 13,
		"?": 14,
		"(": 15,
		";": 16,
		":": 17,
		"<": 18,
		"=": 19,
		"9": 20,
		"3": 21,
		"4": 22,
		"#": 23,
	}

	dec = map[int]string{
		0:  "C",
		1:  "G",
		2:  "A",
		3:  "T",
		4:  "N",
		5:  "@",
		6:  "D",
		7:  "F",
		8:  "H",
		9:  "I",
		10: "J",
		11: "B",
		12: "E",
		13: ">",
		14: "?",
		15: "(",
		16: ";",
		17: ":",
		18: "<",
		19: "=",
		20: "9",
		21: "3",
		22: "4",
		23: "#",
	}
)

func main() {
	s1 := LwzEncoder(seq1, enc)
	s2 := LwzEncoder(seq2, enc)
	s3 := LwzEncoder(mix, enc)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(LwzDecoder(s1, dec))
	fmt.Println(LwzDecoder(s2, dec))
	fmt.Printf("Original string: %d, encoded: %d\n", len(mix), len(s3))
	fmt.Println(s3)
	b := []byte(seq1)
	fmt.Println(b)
}

func LwzEncoder(x string, encoding map[string]int) []int {

	var out []int

	p := ""
	for _, s := range x {
		current := p + string(s)
		if _, ok := encoding[current]; ok {
			p = current
		} else {
			encoding[current] = len(encoding)
			out = append(out, encoding[p])
			p = string(s)
		}
	}
	// when string is empty append the code for prefix
	out = append(out, encoding[p])

	return out
}

func LwzDecoder(x []int, encoding map[int]string) string {
	decoded := encoding[x[0]]
	oldcode := x[0]
	p := ""
	for _, code := range x[1:] {
		if val, ok := encoding[code]; ok {
			decoded += val
			p = encoding[oldcode]
			k := string(val[0])
			encoding[len(encoding)] = p + k
			oldcode = code
		} else {
			p = encoding[oldcode]
			k := string(p[0])
			decoded += p + k
			encoding[len(encoding)] = p + k
			oldcode = code
		}
	}
	return decoded
}

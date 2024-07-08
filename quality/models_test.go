package quality

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name  string
		input DNA
		want  bool
	}{
		{
			name:  "Return false when DNA containse bases other than A, C, T, G, N",
			input: "ACTGNCTGGGGNCTGGGCNTAAAAATYR",
			want:  false,
		},
		{
			name:  "Return true when DNA contains ACTGN only",
			input: "GGNAGGAAATATTGNAGAGGATCTCCATAACNATCTCCTGAAGAAGAGGTAACCGGGTATTCCACACACCCGGAATGTTGTGCACTCACTCTAATTTTCAAAAGTAATCACAACAAGATTATTTATGATNAAATGTGACGCGACNCGCAA",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dna := tt.input
			res := dna.Validate()
			if res != tt.want {
				t.Fatal("should return false for")
			}
		})
	}
}

func TestCounter(t *testing.T) {
	tests := []struct {
		name  string
		input DNA
		want  map[string]int64
	}{
		{
			name:  "Return false when DNA containse bases other than A, C, T, G, N",
			input: "GGAAANAGGAAATATTGNAGAGGATCTCCATAACNATCTCCTGAAGAAGAGGTAACCGGGTATTCCACACACCCGGAATGTTGTGCACTCACTCTAATTTTCAAAAGTAATCACAACAAGATTATTTATGATNAAATGTGACGCGACNCGCAA",
			want:  map[string]int64{"A": 51, "C": 30, "G": 29, "N": 5, "T": 35},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var counter Counter
			dna := tt.input

			var wg sync.WaitGroup

			for _, val := range dna {
				wg.Add(1)

				go func() {
					defer wg.Done()
					go counter.Add(string(val), 1)
				}()

				wg.Wait()

			}

			res := make(map[string]any)

			c, _ := counter.m.Load("C")
			g, _ := counter.m.Load("G")
			u, _ := counter.m.Load("T")
			a, _ := counter.m.Load("A")
			n, _ := counter.m.Load("N")

			res["C"] = c
			res["G"] = g
			res["T"] = u
			res["A"] = a
			res["N"] = n

			if res == nil {
				t.Fatalf("expected %+v, received %+v", tt.want, res)
			}

			str, _ := json.Marshal(res)

			fmt.Printf("The counter: %s\n", str)
		})
	}
}

func TestBaseFrequency(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[rune]int
	}{
		{
			name:  "Returns empty map",
			input: "",
			want:  map[rune]int{},
		},
		{
			name:  "Returns map with frequencies of bases",
			input: "ANNN",
			want:  map[rune]int{65: 1, 78: 3},
		},
		{
			name:  "Returns map with frequencies of bases",
			input: "GGNAGGAAATATTGNAGAGGATCTCCATAACNATCTCCTGAAGAAGAGGTAACCGGGTATTCCACACACCCGGAATGTTGTGCACTCACTCTAATTTTCAAAAGTAATCACAACAAGATTATTTATGATNAAATGTGACGCGACNCGCAA",
			want:  map[rune]int{65: 51, 67: 30, 71: 29, 84: 35, 78: 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dna := tt.input
			res := BaseFrequency(dna)
			if res[65] != tt.want[65] {
				t.Fatalf("wrong count for A: expect %d, got %d", tt.want[65], res[65])
			}
			if res[67] != tt.want[67] {
				t.Fatalf("wrong count for C: expect %d, got %d", tt.want[67], res[67])
			}
			if res[71] != tt.want[71] {
				t.Fatalf("wrong count for G: expect %d, got %d", tt.want[71], res[71])
			}
			if res[84] != tt.want[84] {
				t.Fatalf("wrong count for T: expect %d, got %d", tt.want[84], res[84])
			}
			if res[78] != tt.want[78] {
				t.Fatalf("wrong count for N: expect %d, got %d", tt.want[78], res[78])
			}
		})
	}
}

func TestParseHeder(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]int
	}{
		{
			name:  "Returns correct header",
			input: "@M07197:20:000000000-K7JRN:1:1101:16242:1028 1:N:0:8",
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := ParseHeader(tt.input)
			fmt.Printf("Received header: %#v\n", res)

		})
	}
}

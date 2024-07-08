package util

import "testing"

const (
	header_valid    = `@ERR101899.1 M10_151:1:2:13999:1320 length=150`
	dna_valid_upper = `ACCTGGCTNGACAAAAAAAAAANAAAAAATTTGTTTGTNTTACTG`
	dna_valid_lower = `acctggtgttgcnaaataagtagtaaanaattttggggagattgagatagacagatagacag`
	dna_valid_mixed = `ACCTGGCNTGACaagtagtaaaaanttttggggagAAAAAAAA`
	dna_invalid     = `aaacctgatgcnctgacraaactg`
	util_string     = `+ERR101899.5 M10_151:1:2:15495:1326 length=150`
	quality_valid   = `@<@DFFADFFFHHGI@HIIHDIFCDGGIIFIGG?:EHIIJJIIEGEIGAEGGGFEEHEIIG:CHGHIIIAHGC;)7@@ECEHB@BF@C(66>ADDDD?AC@C>5<CC@C@DCCDDDCDA:CCA@CCDC>C?A>A::>+>8?33>@A>A##`
	quality_invalid = `@<@DFFADFFFHHGI@HIIHDIFC DGGIIFIGG?:EHIIJJIIEGEIGAEGGGFEEHEIIG:CHGHIIIAHGC;)7@@ECEHB@BF@C(66>ADDDD?AC@C>5<CC@C@DCCDDDCDA:CCA@CCDC>C?A>A::>+>8?33>@A>A##`
)

func TestIsDNA(t *testing.T) {
	t.Run("Test that valid dna string (upper case) is matched correctly", func(t *testing.T) {
		got := isDNA(dna_valid_upper)
		want := true
		if got != want {
			t.Errorf("Want %t, got %t\n", want, got)
		}
	})

	t.Run("Test that valid dna string (lower case) is matched correctly", func(t *testing.T) {
		got := isDNA(dna_valid_lower)
		want := true
		if got != want {
			t.Errorf("Want %t, got %t\n", want, got)
		}
	})

	t.Run("Test that valid dna string (mixed case) is matched correctly", func(t *testing.T) {
		got := isDNA(dna_valid_mixed)
		want := true
		if got != want {
			t.Errorf("Want %t, got %t\n", want, got)
		}
	})

	t.Run("Test that invalid dna string is detected", func(t *testing.T) {
		got := isDNA(dna_invalid)
		want := false
		if got != want {
			t.Errorf("Want %t, got %t\n", want, got)
		}
	})
}

func TestIsHeader(t *testing.T) {
	t.Run("Test that valid header (upper case) is matched correctly", func(t *testing.T) {
		got := isHeader(header_valid)
		want := true
		if got != want {
			t.Errorf("Want %t, got %t\n", want, got)
		}
	})

	t.Run("Test that invalid header is detected", func(t *testing.T) {
		got := isHeader(quality_valid)
		want := false
		if got != want {
			t.Errorf("Want %t, got %t\n", want, got)
		}
	})
}

func TestIsUtilString(t *testing.T) {
	t.Run("Test that util string is detected", func(t *testing.T) {
		got := isUtilString(util_string)
		want := true
		if got != want {
			t.Errorf("Want %t, got %t\n", want, got)
		}
	})

	t.Run("Test that empty util string is detected", func(t *testing.T) {
		got := isUtilString("+")
		want := true
		if got != want {
			t.Errorf("Want %t, got %t\n", want, got)
		}
	})

	t.Run("Test that not util string returns false", func(t *testing.T) {
		got := isUtilString(dna_invalid)
		want := false
		if got != want {
			t.Errorf("Want %t, got %t\n", want, got)
		}
	})
}

func TestIsQualityString(t *testing.T) {
	t.Run("Test that valid quality string is detected", func(t *testing.T) {
		got := isQualityString(quality_valid)
		want := true
		if got != want {
			t.Errorf("Want %t, got %t\n", want, got)
		}
	})

	t.Run("Test that invalid quality string is detected", func(t *testing.T) {
		got := isQualityString(quality_invalid)
		want := false
		if got != want {
			t.Errorf("Want %t, got %t\n", want, got)
		}
	})
}

func TestIdentifyReadLine(t *testing.T) {
	t.Run("Test dna can be detected", func(t *testing.T) {
		got := IdentifyReadLine(dna_valid_upper)
		want := "dna"
		if got != want {
			t.Errorf("Want %s, got %s\n", want, got)
		}
	})

	t.Run("Test invalid dna returns unknown", func(t *testing.T) {
		got := IdentifyReadLine(dna_invalid)
		want := "unknown"
		if got != want {
			t.Errorf("Want %s, got %s\n", want, got)
		}
	})
	t.Run("Test quality string can be detected", func(t *testing.T) {
		got := IdentifyReadLine(quality_valid)
		want := "quality"
		if got != want {
			t.Errorf("Want %s, got %s\n", want, got)
		}
	})

	t.Run("Test invalid quality string returns unknown", func(t *testing.T) {
		got := IdentifyReadLine(quality_invalid)
		want := "unknown"
		if got != want {
			t.Errorf("Want %s, got %s\n", want, got)
		}
	})
}

package rnatranscription

func ToRNA(dna string) string {
	dnaToRna := map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}
	dnaRunes := []rune(dna)
	rna := ""

	for _, l := range dnaRunes {
		value, ok := dnaToRna[string(l)]
		if !ok {
			panic("Value didn't found")
		}
		rna += value
	}

	return rna
}

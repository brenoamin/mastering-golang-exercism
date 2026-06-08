package proteintranslation

import (
    "errors"
)

var ErrStop = errors.New("stop immediately, the codon is already complete")

var ErrInvalidBase = errors.New("invalid Base")

func FromRNA(rna string) ([]string, error) {

	var result []string
	for i := 0; i <= len(rna)-1; i += 3 {
        if i+3 > len(rna) {
            return nil, ErrInvalidBase
        }
		val, err := FromCodon(rna[i : i+3])
		if errors.Is(err, ErrStop) {
			return result, nil
		}
		if val == "" {
			return result, ErrInvalidBase
		}
		result = append(result, val)

	}
	return result, nil
}

func FromCodon(codon string) (string, error) {
	codonToAminoAcid := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}

	val, ok := codonToAminoAcid[codon]

	if !ok {
		return "", ErrInvalidBase
	}
	if val == "STOP" {
		return "", ErrStop
	}
	return val, nil
}


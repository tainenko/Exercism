package protein

import "errors"

var codon2Protein = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP"}

// ErrStop is returned when we meet STOP codon
var ErrStop = errors.New("codon stop code")

// ErrInvalidBase returned when we meet Invalid codon
var ErrInvalidBase = errors.New("invalid base")

//FromRNA converts codons to protein list
func FromRNA(rna string) ([]string, error) {
	var proteins []string
	for i := 0; i < len(rna); i += 3 {
		nucleotide := rna[i : i+3]
		protein, err := FromCodon(nucleotide)
		if err == ErrStop {
			return proteins, nil
		} else if err == ErrInvalidBase {
			return proteins, ErrInvalidBase
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}

// FromCodon converts a codon to a protein
func FromCodon(codon string) (string, error) {
	protein := codon2Protein[codon]
	if protein == "STOP" {
		return "", ErrStop
	}
	if protein == "" {
		return "", ErrInvalidBase
	}
	return protein, nil
}

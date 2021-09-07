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

var ErrStop = errors.New("codon stop code")
var ErrInvalidBase = errors.New("invalid base")

func FromRNA(rna string) (codons []string, err error) {
	codon := ""
	for i := 0; i < len(rna); i += 3 {
		nucleotide := rna[i : i+3]
		codon, err = FromCodon(nucleotide)
		if err == ErrStop {
			return codons, nil
		} else if err == ErrInvalidBase {
			return codons, ErrInvalidBase
		}
		codons = append(codons, codon)
	}
	return codons, nil
}
func FromCodon(codon string) (nucleotide string, err error) {
	protein := codon2Protein[codon]
	if protein == "STOP" {
		return "", ErrStop
	}
	if protein == "" {
		return "", ErrInvalidBase
	}
	return protein, nil
}

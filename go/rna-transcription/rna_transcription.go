package strand

var toRnaMap = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA convert input dna string and return a rna string
func ToRNA(dna string) string {
	rna := make([]rune, len(dna))
	for k, r := range dna {
		rna[k] = toRnaMap[r]
	}
	return string(rna)

}

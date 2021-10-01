package rotationalcipher

// RotationalCipher would shift the input plain and return
func RotationalCipher(plain string, shiftKey int) string {
	shiftKey %= 26
	if shiftKey == 0 {
		return plain
	}
	var shifted []rune
	var newChr rune
	for _, chr := range plain {
		if chr >= 'a' && chr <= 'z' {
			newChr = rune((int(chr)-'a'+shiftKey)%26) + 'a'
		} else if chr >= 'A' && chr <= 'Z' {
			newChr = rune((int(chr)-'A'+shiftKey)%26) + 'A'
		} else {
			newChr = chr
		}
		shifted = append(shifted, newChr)
	}
	return string(shifted)

}

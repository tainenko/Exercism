package series

// All  returns a list of all substrings of s with length n
func All(n int, s string) []string {
	var out []string
	for i := 0; i < len(s)-n+1; i++ {
		out = append(out, s[i:i+n])
	}
	return out
}

// UnsafeFirst returns the first substring of s with length n
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

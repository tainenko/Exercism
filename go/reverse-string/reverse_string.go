package reverse

// Reverse would return a reversed string of input
func Reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

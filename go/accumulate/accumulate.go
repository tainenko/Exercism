package accumulate

func Accumulate(strs []string, operator func(string) string) (result []string) {
	for i := range strs {
		result = append(result, operator(strs[i]))
	}
	return
}

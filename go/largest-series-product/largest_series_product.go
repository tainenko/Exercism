package lsproduct

import "errors"

// LargestSeriesProduct would return the largest series n product in given string
func LargestSeriesProduct(s string, n int) (int, error) {
	if len(s) < n {
		return -1, errors.New("string is shorter than span")
	}

	if n < 0 {
		return -1, errors.New("negative span")
	}

	if n == 0 {
		return 1, nil
	}

	tmp := 1
	cnt := 0
	var result int
	for i := 0; i < len(s); {
		if s[i] < '0' || s[i] > '9' {
			return -1, errors.New("input contains alpha-characters")
		}

		if s[i] == '0' {
			i++
			tmp = 1
			cnt = 0
			continue
		}
		tmp *= int(s[i] - '0')
		cnt++
		if cnt > n {
			tmp /= int(s[i-n] - '0')
			cnt--
		}
		result = max(result, tmp)
		i++
	}
	if cnt == n {
		return result, nil
	}
	return 0, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

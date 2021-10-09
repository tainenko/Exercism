package wordy

import (
	"regexp"
	"strconv"
)

// Answer parse and evaluate simple math word problems returning the answer as an integer.
func Answer(s string) (int, bool) {
	if match, _ := regexp.MatchString(`What is -?\d+(?: (?:plus|minus|multiplied by|divided by) -?\d+)*\?`, s); !match {
		return 0, false
	}

	reg1 := regexp.MustCompile(`-?\d+`)
	nums := reg1.FindAllString(s, -1)

	reg2 := regexp.MustCompile(`plus|minus|multiplied by|divided by`)
	ops := reg2.FindAllString(s, -1)

	if len(nums) < len(ops)-1 {
		return 0, false
	}
	sum, _ := strconv.Atoi(nums[0])
	for i := range ops {
		n, _ := strconv.Atoi(nums[i+1])
		switch ops[i] {
		case "plus":
			sum += n
		case "minus":
			sum -= n
		case "multiplied by":
			sum *= n
		case "divided by":
			sum /= n
		}
	}
	return sum, true
}

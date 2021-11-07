package encode

import (
	"strconv"
	"strings"
)

// RunLengthEncode replace the input string by just one data value and count.
func RunLengthEncode(input string) string {
	builder := strings.Builder{}
	if len(input) == 0 {
		return input
	}
	curr := input[0]
	cnt := 1
	for i := 1; i < len(input); i++ {
		if curr == input[i] {
			cnt++
		} else {
			if cnt != 1 {
				builder.WriteString(strconv.Itoa(cnt))
			}
			builder.WriteByte(curr)
			curr = input[i]
			cnt = 1
		}
	}
	if cnt != 1 {
		builder.WriteString(strconv.Itoa(cnt))
	}
	builder.WriteByte(curr)
	return builder.String()
}

// RunLengthDecode decodes the input string from one data value and count to string with repeated letters.
func RunLengthDecode(input string) string {
	builder := strings.Builder{}
	tmp := 0
	for i := 0; i < len(input); i++ {

		if num, err := strconv.Atoi(string(input[i])); err == nil {
			tmp = tmp*10 + num
			continue
		}
		if tmp == 0 {
			builder.WriteByte(input[i])
			continue
		}
		for j := 0; j < tmp; j++ {
			builder.WriteByte(input[i])
		}
		tmp = 0
	}
	return builder.String()
}

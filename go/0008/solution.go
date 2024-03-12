package problem0008

import (
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	const (
		lowBound  = -(1 << 31)
		highBound = (1 << 31) - 1
	)
	var (
		trimmed = strings.TrimSpace(s)
		sign    = 1
		result  = 0
	)

	switch trimmed[0] {
	case '-':
		sign = -1
		trimmed = trimmed[1:]
	case '+':
		trimmed = trimmed[1:]
	}

	for _, char := range trimmed {
		if !unicode.IsDigit(char) {
			break
		}

		num := int(char - '0')
		result = result*10 + num
		if result > highBound {
			if sign < 0 {
				return lowBound
			}
			return highBound
		}
	}

	result *= sign

	return result
}

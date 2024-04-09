package problem0013

func romanToInt(str string) int {
	symbols := []string{
		"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX",
		"V", "IV", "I",
	}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := 0

	for idx, sym := range symbols {
		if len(str) == 0 {
			break
		}

		length := len(sym)

		for len(str) >= length && str[:length] == sym {
			result += values[idx]
			str = str[length:]
		}
	}

	return result
}

package problem0012

func intToRoman(num int) string {
	symbols := []string{
		"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX",
		"V", "IV", "I",
	}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""

	for idx, val := range values {
		for num >= val {
			result += symbols[idx]
			num -= val
		}
	}

	return result
}

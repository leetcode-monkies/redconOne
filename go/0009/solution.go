package problem0009

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	result, num := 0, x

	for num != 0 {
		result = result*10 + num%10
		num = num / 10
	}

	return x == result
}

package problem0007

func reverse(x int) int {
	sign := 1
	result := 0
	lowBound := -(1 << 31)
	highBound := (1 << 31) - 1
	if x < 0 {
		sign = -1
	}
	x *= sign

	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}

	result *= sign

	if result < lowBound || result > highBound {
		return 0
	}

	return result
}

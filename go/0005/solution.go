package problem0005

func longestPalindrome(s string) string {
	for length := len(s) - 1; length >= 0; length-- {
		for start := 0; start+length < len(s); start++ {
			if isPalindrome(start, start+length, s) {
				return s[start : start+length+1]
			}
		}
	}

	return ""
}

func isPalindrome(left, right int, s string) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

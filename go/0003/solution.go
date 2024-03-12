package problem0003

func lengthOfLongestSubstring(s string) int {
	dict := make(map[byte]int, 0)
	var longest int = 0

	for p1, p2 := 0, 0; p2 < len(s); {
		if _, ok := dict[s[p2]]; !ok || p1 == p2 {
			dict[s[p2]] = 1
			p2++
		} else {
			delete(dict, s[p1])
			p1++
		}
		if longest < p2-p1 {
			longest = p2 - p1
		}
	}

	return longest
}

package problem0791

import "strings"

func CustomSortString(order, s string) string {
	result := strings.Builder{}
	freq := [26]int{}

	for _, r := range s {
		freq[r-'a']++
	}

	for _, r := range order {
		count := freq[r-'a']
		if count == 0 {
			continue
		}
		freq[r-'a'] = 0
		result.WriteString(strings.Repeat(string(r), count))
	}

	for r, count := range freq {
		if count > 0 {
			result.WriteString(strings.Repeat(string(byte(r+'a')), count))
		}
	}

	return result.String()
}

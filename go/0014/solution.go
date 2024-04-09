package problem0014

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for _, word := range strs {
		i := 0

		for ; i < len(word) && i < len(prefix) && word[i] == prefix[i]; i++ {
		}

		prefix = prefix[:i]
	}

	return prefix
}

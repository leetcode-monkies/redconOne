package problem0003

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testcases := []struct {
		s    string
		want int
	}{
		{
			s:    "abcabcbb",
			want: 3,
		},
		{
			s:    "bbbbb",
			want: 1,
		},
		{
			s:    "pwwkew",
			want: 3,
		},
	}

	for _, tc := range testcases {
		result := lengthOfLongestSubstring(tc.s)
		if result != tc.want {
			t.Errorf("lengthOfLongestSubstring: %v, want: %v\n", result, tc.want)
		}
	}
}

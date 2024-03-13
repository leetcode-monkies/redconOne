package problem0005

import "testing"

func TestLongestPalindrome(t *testing.T) {
	testcases := []struct {
		s, want string
	}{
		{
			s:    "babad",
			want: "bab",
		},
		{
			s:    "cbbd",
			want: "bb",
		},
		{
			s:    "a",
			want: "a",
		},
	}

	for _, tc := range testcases {
		result := longestPalindrome(tc.s)
		if result != tc.want {
			t.Errorf("longestPalindrome: %v, want: %v\n", result, tc.want)
		}
	}
}

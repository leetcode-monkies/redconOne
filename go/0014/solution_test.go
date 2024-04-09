package problem0014

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	testcases := []struct {
		want string
		strs []string
	}{
		{
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			strs: []string{
				"dog", "racecar", "car",
			},
			want: "",
		},
	}

	for _, tc := range testcases {
		result := longestCommonPrefix(tc.strs)
		if result != tc.want {
			t.Errorf("longestCommonPrefix: %v, want: %v", result, tc.want)
		}
	}
}

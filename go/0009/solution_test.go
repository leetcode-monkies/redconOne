package problem0009

import "testing"

func TestIsPalindrome(t *testing.T) {
	testcases := []struct {
		x    int
		want bool
	}{
		{
			x:    121,
			want: true,
		},
		{
			x:    -121,
			want: false,
		},
		{
			x:    10,
			want: false,
		},
	}

	for _, tc := range testcases {
		result := isPalindrome(tc.x)
		if result != tc.want {
			t.Errorf("isPalindrome: %v, want: %v\n", result, tc.want)
		}
	}
}

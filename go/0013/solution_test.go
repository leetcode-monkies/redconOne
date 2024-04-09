package problem0013

import "testing"

func TestRomanToInt(t *testing.T) {
	testcases := []struct {
		str  string
		want int
	}{
		{
			str:  "III",
			want: 3,
		},
		{
			str:  "LVIII",
			want: 58,
		},
		{
			str:  "MCMXCIV",
			want: 1994,
		},
		{
			str:  "DCXXI",
			want: 621,
		},
	}

	for _, tc := range testcases {
		result := romanToInt(tc.str)
		if result != tc.want {
			t.Errorf("romanToInt: %v, want: %v", result, tc.want)
		}
	}
}

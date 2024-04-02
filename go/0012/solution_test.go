package problem0012

import "testing"

func TestIntToRoman(t *testing.T) {
	testcases := []struct {
		want string
		num  int
	}{
		{
			num:  3,
			want: "III",
		},
		{
			num:  58,
			want: "LVIII",
		},
		{
			num:  1994,
			want: "MCMXCIV",
		},
	}

	for _, tc := range testcases {
		result := intToRoman(tc.num)
		if result != tc.want {
			t.Errorf("intToRoman: %v, want: %v", result, tc.want)
		}
	}
}

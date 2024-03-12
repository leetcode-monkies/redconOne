package problem0008

import "testing"

func TestMyAtoi(t *testing.T) {
	testcases := []struct {
		s    string
		want int
	}{
		{
			s:    "42",
			want: 42,
		},
		{
			s:    "       -42",
			want: -42,
		},
		{
			s:    "4193 with words",
			want: 4193,
		},
	}

	for _, tc := range testcases {
		result := myAtoi(tc.s)
		if result != tc.want {
			t.Errorf("myAtoi: %v, want: %v", result, tc.want)
		}
	}
}

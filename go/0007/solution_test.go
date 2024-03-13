package problem0007

import "testing"

func TestReverse(t *testing.T) {
	testcases := []struct {
		x, want int
	}{
		{
			x:    123,
			want: 321,
		},
		{
			x:    -123,
			want: -321,
		},
		{
			x:    120,
			want: 21,
		},
	}

	for _, tc := range testcases {
		result := reverse(tc.x)
		if result != tc.want {
			t.Errorf("reverse: %v, want: %v", result, tc.want)
		}

	}
}

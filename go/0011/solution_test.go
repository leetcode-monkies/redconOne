package problem0011

import "testing"

func TestMaxArea(t *testing.T) {
	testcases := []struct {
		height []int
		want   int
	}{
		{
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			height: []int{1, 1},
			want:   1,
		},
	}

	for _, tc := range testcases {
		result := maxArea(tc.height)
		if result != tc.want {
			t.Errorf("maxArea: %v, want: %v\n", result, tc.want)
		}
	}
}

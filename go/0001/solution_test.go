package problem0001

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testcases := []struct {
		nums, want []int
		target     int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

	for _, tc := range testcases {
		result := TwoSum(tc.nums, tc.target)
		if !reflect.DeepEqual(result, tc.want) {
			t.Errorf("TwoSum: %q, want: %q\n", result, tc.want)
		}
	}
}

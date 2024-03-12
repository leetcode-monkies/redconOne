package problem0791

import (
	"testing"
)

func TestCustomSortString(t *testing.T) {
	testcases := []struct {
		order, s, want string
	}{
		{
			order: "cba",
			s:     "abcd",
			want:  "cbad",
		},
		{
			order: "bcafg",
			s:     "abcd",
			want:  "bcad",
		},
	}

	for _, testcase := range testcases {
		result := CustomSortString(testcase.order, testcase.s)
		if result != testcase.want {
			t.Errorf("CustomSortString: %q, want: %q\n", result, testcase.want)
		}
	}
}

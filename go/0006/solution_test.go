package problem0006

import "testing"

func TestConvert(t *testing.T) {
	testcases := []struct {
		s, want string
		numRows int
	}{
		{
			s:       "PAYPALISHIRING",
			numRows: 3,
			want:    "PAHNAPLSIIGYIR",
		},
		{
			s:       "PAYPALISHIRING",
			numRows: 4,
			want:    "PINALSIGYAHRPI",
		},
		{
			s:       "A",
			numRows: 1,
			want:    "A",
		},
	}

	for _, tc := range testcases {
		result := convert(tc.s, tc.numRows)
		if result != tc.want {
			t.Errorf("convert: %v, want: %v", result, tc.want)
		}
	}
}

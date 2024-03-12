package problem0006

import "strings"

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	rows := make([]string, numRows)
	direction := 1
	row := 0

	for _, rune := range s {
		rows[row] += string(rune)
		row += direction
		if row == 0 || row == numRows-1 {
			direction *= -1
		}
	}

	return strings.Join(rows, "")
}

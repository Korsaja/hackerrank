package medium

import (
	"math"
	"strings"
)

const paddingByte = '0'

func encryption(s string) string {
	rows, cols := getRowsAndColumn(s)
	matrix := buildMatrix(rows, cols, s)

	var out strings.Builder
	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			v := matrix[r][c]
			if v != paddingByte {
				out.WriteRune(v)
			}
		}
		if c < cols-1 {
			out.WriteRune(' ')
		}
	}
	return out.String()
}

func buildMatrix(rows, cols int, s string) [][]rune {
	matrix := make([][]rune, rows)
	text := []rune(s)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]rune, cols)
		n := copy(matrix[i], text)
		text = text[n:]
		if n < cols {
			for n < cols {
				// padding
				matrix[i][n] = paddingByte
				n++
			}
		}
	}
	return matrix
}

func getRowsAndColumn(s string) (int, int) {
	l := len(s)
	sqrt := math.Sqrt(float64(l))
	rows := int(math.Floor(sqrt))
	cols := int(math.Ceil(sqrt))

	if rows*cols < l {
		return cols, cols
	}

	return rows, cols
}

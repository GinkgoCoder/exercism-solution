package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	rowView [][]int
	colView [][]int
}

func New(s string) (Matrix, error) {
	if s == "" {
		return Matrix{}, nil
	}
	rows := strings.Split(s, "\n")
	rowNum, colNum := len(rows), len(strings.Fields(rows[0]))
	matrix := Matrix{
		rowView: make([][]int, rowNum),
		colView: make([][]int, colNum),
	}
	for i, row := range rows {
		if row == "" {
			return Matrix{}, errors.New("cannot be empty")
		}
		for j, col := range strings.Fields(row) {
			if j >= colNum {
				return Matrix{}, errors.New("column number is not correct")
			}

			if num, err := strconv.Atoi(col); err != nil {
				return Matrix{}, err
			} else {
				matrix.rowView[i] = append(matrix.rowView[i], num)
				matrix.colView[j] = append(matrix.colView[j], num)
			}
		}
	}
	return matrix, nil
}
func copySlice2D(original [][]int) [][]int {
	duplicate := make([][]int, len(original))

	for i := range original {
		duplicate[i] = append([]int(nil), original[i]...)
	}

	return duplicate
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	return copySlice2D(m.colView)
}

func (m Matrix) Rows() [][]int {
	return copySlice2D(m.rowView)
}

func (m Matrix) Set(row, col, val int) bool {
	if row >= 0 && row < len(m.rowView) && col >= 0 && col < len(m.colView) {
		m.rowView[row][col] = val
		m.colView[col][row] = val
		return true
	}
	return false
}

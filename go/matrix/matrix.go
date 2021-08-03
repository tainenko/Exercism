package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix is a 2D integer array
type Matrix [][]int

// New function receive a string as input and return a integer matrix
func New(s string) (Matrix, error) {
	rows := strings.Split(s, "\n")
	n := len(rows)
	m := make(Matrix, n)
	var err error
	var rowLength int
	for i, row := range rows {
		cells := strings.Split(strings.TrimSpace(row), " ")
		if i == 0 {
			rowLength = len(cells)
		} else if rowLength != len(cells) {
			return nil, errors.New("row length must be constant")
		}

		m[i] = make([]int, len(cells))
		for j, cell := range cells {
			if m[i][j], err = strconv.Atoi(cell); err != nil {
				return nil, err
			}
		}
	}
	return m, nil
}

// Rows method would return a copy of rows of matrix
func (m Matrix) Rows() [][]int {
	newMatrix := make([][]int, len(m))
	for i, row := range m {
		newMatrix[i] = make([]int, len(row))
		copy(newMatrix[i], row)
	}
	return newMatrix
}

// Cols method would return a copy of cols of matrix
func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))
	for col := range m[0] {
		for row := range m {
			cols[col] = append(cols[col], m[row][col])
		}
	}
	return cols
}

// Set function could set a new val to matrix[r][c]
func (m Matrix) Set(r, c, val int) bool {
	if r < 0 || r >= len(m) || c < 0 || c >= len(m[0]) {
		return false
	}
	m[r][c] = val
	return true
}

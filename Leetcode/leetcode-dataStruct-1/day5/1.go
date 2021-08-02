package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rows, columns := 9, 9

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			d := board[r][c]
			if d != '.' {
				if !validDigit(d) || inRow(board[r], c, d) || inColumn(board, c, r, d) || inRegion(board, r, c, d) {
					return false
				}
			}

		}
	}
	return true
}

// todo

func main() {
	// fmt.Println("1:", '1', "9:", '9')
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	res := isValidSudoku(board)
	fmt.Println(res)
}

func validDigit(digit byte) bool {
	return digit >= '1' && digit <= '9'
}

func inRow(rows []byte, col int, digit byte) bool {
	for c := 0; c < 9; c++ {
		if rows[c] == digit && c != col {
			return true
		}
	}
	return false
}

func inColumn(g [][]byte, column, row int, digit byte) bool {
	for r := 0; r < 9; r++ {
		if g[r][column] == digit && r != row {
			return true
		}
	}
	return false
}

func inRegion(g [][]byte, row, column int, digit byte) bool {
	startRow, startColumn := row/3*3, column/3*3
	for r := startRow; r < startRow+3; r++ {
		for c := startColumn; c < startColumn+3; c++ {
			if g[r][c] == digit && r != row && c != column {
				return true
			}
		}
	}
	return false
}

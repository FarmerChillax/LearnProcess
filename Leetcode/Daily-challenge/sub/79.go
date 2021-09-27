package sub

func exist(board [][]byte, word string) bool {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			res := dfs(board, word, col, row)
			if res {
				return true
			}
		}
	}
	return false
}

func Sub79(board [][]byte, word string) bool { return exist(board, word) }

func dfs(board [][]byte, word string, col, row int) bool {
	if len(word) < 1 {
		return true
	}
	if col < 0 || row < 0 || row == len(board) || col == len(board[row]) {
		return false
	}
	if board[row][col] != word[0] {
		return false
	}

	board[row][col] ^= 255
	exit := dfs(board, word[1:], col-1, row) ||
		dfs(board, word[1:], col+1, row) ||
		dfs(board, word[1:], col, row-1) ||
		dfs(board, word[1:], col, row+1)
	board[row][col] ^= 255
	return exit
}

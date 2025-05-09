package leetcode

// Given an m x n grid of characters board and a string word, return true if word exists in the grid.

// The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	var dfs func(i, j, k int) bool

	dfs = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}
		if i < 0 || i >= rows || j < 0 || j >= cols || board[i][j] != word[k] {
			return false
		}

		temp := board[i][j]
		board[i][j] = '#'

		found := dfs(i+1, j, k+1) ||
			dfs(i-1, j, k+1) ||
			dfs(i, j+1, k+1) ||
			dfs(i, j-1, k+1)

		board[i][j] = temp

		return found
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

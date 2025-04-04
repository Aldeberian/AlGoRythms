package leetcode

// Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.

func minPathSum(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if j > 0 && i > 0 {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			} else if j > 0 {
				grid[i][j] += grid[i][j-1]
			} else if i > 0 {
				grid[i][j] += grid[i-1][j]
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

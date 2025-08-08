package leetcode

import "fmt"

// There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]).
// The robot can only move either down or right at any point in time.

// Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

// The test cases are generated so that the answer will be less than or equal to 2 * 109.

func uniquePaths(m int, n int) int {
	dp := make(map[string]int)
	var rec func(int, int) int

	rec = func(i, j int) int {
		s := fmt.Sprintf("%d,%d", i, j)
		if val, ok := dp[s]; ok {
			return val
		}
		if i == m-1 && j == n-1 {
			return 1
		}

		right, bottom := 0, 0
		if j < n-1 {
			right = rec(i, j+1)
		}
		if i < m-1 {
			bottom = rec(i+1, j)
		}

		dp[s] = right + bottom

		return right + bottom
	}

	return rec(0, 0)
}

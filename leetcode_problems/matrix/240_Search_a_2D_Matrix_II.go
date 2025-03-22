package leetcode

import (
	"fmt"
	"strconv"
)

// Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the following properties:

//     Integers in each row are sorted in ascending from left to right.
//     Integers in each column are sorted in ascending from top to bottom.

func SearchMatrix(matrix [][]int, target int) bool {
	i, j := 0, 0
	visited := make(map[string]bool)

	for {
		fmt.Printf("i : %d, j : %d, val : %d\n", i, j, matrix[i][j])
		visited[strconv.Itoa(i)+","+strconv.Itoa(j)] = true
		if matrix[i][j] == target {
			return true
		}

		if j < len(matrix[0])-1 && matrix[i][j+1] <= target && !visited[strconv.Itoa(i)+","+strconv.Itoa(j+1)] {
			j++
		} else if i < len(matrix)-1 && matrix[i+1][j] <= target && !visited[strconv.Itoa(i+1)+","+strconv.Itoa(j)] {
			i++
		} else if i > 0 && matrix[i-1][j] <= target && !visited[strconv.Itoa(i-1)+","+strconv.Itoa(j)] {
			i--
		} else if j > 0 && matrix[i][j-1] <= target && !visited[strconv.Itoa(i)+","+strconv.Itoa(j-1)] {
			j--
		} else if !visited[strconv.Itoa(i)+","+strconv.Itoa(j-1)] {
			j--
		} else if !visited[strconv.Itoa(i)+","+strconv.Itoa(j+1)] {
			j++
		} else if !visited[strconv.Itoa(i-1)+","+strconv.Itoa(j)] {
			i--
		} else if !visited[strconv.Itoa(i+1)+","+strconv.Itoa(j)] {
			i++
		} else {
			return false
		}
	}
}

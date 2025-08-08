package leetcode

import "math"

// You are given an m x n matrix maze (0-indexed) with empty cells (represented as '.') and walls (represented as '+').
// You are also given the entrance of the maze, where entrance = [entrancerow, entrancecol] denotes the row and column of the cell you are initially standing at.

// In one step, you can move one cell up, down, left, or right. You cannot step into a cell with a wall, and you cannot step outside the maze.
// Your goal is to find the nearest exit from the entrance. An exit is defined as an empty cell that is at the border of the maze. The entrance does not count as an exit.

// Return the number of steps in the shortest path from the entrance to the nearest exit, or -1 if no such path exists.

// Recursivity is not fast enough

func copyMaze(original [][]byte) [][]byte {
	copied := make([][]byte, len(original))
	for i := range original {
		copied[i] = make([]byte, len(original[i]))
		copy(copied[i], original[i])
	}
	return copied
}

func nearestExit(maze [][]byte, entrance []int) int {
	var rec func(maze [][]byte, current []int, first bool) int

	rec = func(maze [][]byte, current []int, first bool) int {
		maze[current[0]][current[1]] = '+'
		mini := math.MaxInt

		if !first && (current[0] == 0 || current[1] == 0 || current[0] == len(maze)-1 || current[1] == len(maze[0])-1) {
			return 0
		}

		if current[0] > 0 && maze[current[0]-1][current[1]] == '.' {
			temp := rec(copyMaze(maze), []int{current[0] - 1, current[1]}, false)
			if temp != -1 {
				mini = min(mini, temp+1)
			}
		}

		if current[1] > 0 && maze[current[0]][current[1]-1] == '.' {
			temp := rec(copyMaze(maze), []int{current[0], current[1] - 1}, false)
			if temp != -1 {
				mini = min(mini, temp+1)
			}
		}

		if current[0] < len(maze)-1 && maze[current[0]+1][current[1]] == '.' {
			temp := rec(copyMaze(maze), []int{current[0] + 1, current[1]}, false)
			if temp != -1 {
				mini = min(mini, temp+1)
			}
		}

		if current[1] < len(maze[0])-1 && maze[current[0]][current[1]+1] == '.' {
			temp := rec(copyMaze(maze), []int{current[0], current[1] + 1}, false)
			if temp != -1 {
				mini = min(mini, temp+1)
			}
		}

		if mini == math.MaxInt {
			return -1
		}

		return mini
	}

	return rec(copyMaze(maze), entrance, true)
}

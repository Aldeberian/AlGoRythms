package leetcode

import "strconv"

// Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

func count(n int) int {
	bits := strconv.FormatInt(int64(n), 2)
	res := 0

	for _, b := range bits {
		if b == '1' {
			res++
		}
	}

	return res
}

func countBits(n int) []int {
	res := []int{}

	for i := range n + 1 {
		res = append(res, count(i))
	}

	return res
}

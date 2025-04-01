package leetcode

import "slices"

// Given an array nums of distinct integers, return all the possible . You can return the answer in any order.

func permute(nums []int) [][]int {
	var response [][]int
	var rec func(current []int)

	rec = func(current []int) {
		if len(current) == len(nums) {
			response = append(response, current)
			return
		}

		for _, val := range nums {
			if !slices.Contains(current, val) {
				temp := append(current, val)
				rec(temp)
			}
		}
	}

	rec([]int{})

	return response
}

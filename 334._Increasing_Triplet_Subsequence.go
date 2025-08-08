package leetcode

import "math"

// Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.

func increasingTriplet(nums []int) bool {
	one := math.MaxInt
	two := math.MaxInt

	for _, val := range nums {
		if val < one {
			one = val
		} else if val < two {
			two = val
		} else {
			return true
		}
	}

	return false
}

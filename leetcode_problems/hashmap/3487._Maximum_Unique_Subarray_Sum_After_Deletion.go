package leetcode

import "math"

// You are given an integer array nums.

// You are allowed to delete any number of elements from nums without making it empty. After performing the deletions, select a

// of nums such that:

//     All elements in the subarray are unique.
//     The sum of the elements in the subarray is maximized.

// Return the maximum sum of such a subarray.

func maxSum(nums []int) int {
	total := 0
	maxi := math.MinInt
	seen := make(map[int]bool)

	for _, val := range nums {
		maxi = max(val, maxi)

		if val > 0 && !seen[val] {
			seen[val] = true
			total += val
		}
	}

	if total == 0 {
		return maxi
	}
	return total
}

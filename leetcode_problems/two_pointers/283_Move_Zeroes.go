package leetcode

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

func moveZeroes(nums []int) {
	count := 0
	index := 0

	for _, val := range nums {
		if val == 0 {
			count++
		} else {
			nums[index] = val
			index++
		}
	}

	for i := len(nums) - count; i < len(nums) && i > 0; i++ {
		nums[i] = 0
	}
}

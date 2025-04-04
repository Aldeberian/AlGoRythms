package leetcode

// Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

// There is only one repeated number in nums, return this repeated number.

// You must solve the problem without modifying the array nums and using only constant extra space.

func findDuplicate(nums []int) int {
	compte := make([]int, len(nums))

	for _, val := range nums {
		compte[val]++
	}

	for i, val := range compte {
		if val > 2 {
			return i
		}
	}

	return 0
}

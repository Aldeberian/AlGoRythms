package leetcode

// You are given an array of integers nums. Return the length of the longest of nums which is either strict increasing or strict decrease.

func longestMonotonicSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max, tempMax, increasing := 1, 1, 0

	for i := 1; i < len(nums); i++ {
		if increasing == 0 {
			if nums[i-1] > nums[i] {
				tempMax++
				increasing = 1
			} else if nums[i-1] < nums[i] {
				tempMax++
				increasing = 2
			}
		} else if increasing == 1 {
			if nums[i-1] > nums[i] {
				tempMax++
			} else if nums[i-1] < nums[i] {
				increasing = 0
				tempMax = 0
			}
		} else {
			if nums[i-1] > nums[i] {
				increasing = 0
				tempMax = 0
			} else if nums[i-1] < nums[i] {
				tempMax++
				increasing = 2
			}
		}
		if tempMax > max {
			max = tempMax
		}
	}

	return max
}

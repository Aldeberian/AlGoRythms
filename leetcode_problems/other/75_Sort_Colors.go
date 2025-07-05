package leetcode

// Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

// We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

// You must solve this problem without using the library's sort function.

func sortColors(nums []int) {
	count := []int{0, 0, 0}

	for _, val := range nums {
		count[val]++
	}

	i := 0
	k := 0
	for k < 3 {
		if count[k] > 0 {
			nums[i] = k
			count[k]--
			i++
		} else {
			k++
		}
	}
}

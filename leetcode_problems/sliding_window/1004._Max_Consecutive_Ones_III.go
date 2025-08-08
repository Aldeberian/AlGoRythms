package leetcode

// Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

func longestOnes(nums []int, k int) int {
	start, end := 0, 0
	count := k
	maxi := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			end = max(end+1, start)
		} else {
			if count > 0 {
				count--
				end = max(end+1, start)
			} else {
				i = i - 1
				if nums[start] == 0 {
					count++
				}
				start++
			}
		}

		maxi = max(maxi, end-start)
	}

	return maxi
}

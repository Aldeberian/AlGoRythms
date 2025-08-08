package leetcode

// You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right.
// You can only see the k numbers in the window. Each time the sliding window moves right by one position.
// Return the max sliding window.

func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}
	maxi := 0

	for i := 0; i < k; i++ {
		maxi = max(maxi, nums[i])
	}

	res := []int{maxi}

	left := 0
	right := k - 1

	for right < len(nums)-1 {

	}
}

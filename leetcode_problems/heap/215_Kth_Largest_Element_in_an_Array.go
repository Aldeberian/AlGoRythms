package leetcode

import "slices"

// Given an integer array nums and an integer k, return the kth largest element in the array.

// Note that it is the kth largest element in the sorted order, not the kth distinct element.

// Can you solve it without sorting?

func findKthLargest(nums []int, k int) int {
	heap := nums[:k]
	slices.Sort(heap)

	for i := k; i < k; i++ {
		if nums[i] > heap[0] {
			heap[0] = nums[i]
			slices.Sort(heap)
		}
	}

	return heap[len(heap)]
}

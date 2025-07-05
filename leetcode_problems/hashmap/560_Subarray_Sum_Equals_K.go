package leetcode

// Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

// A subarray is a contiguous non-empty sequence of elements within an array.

func subarraySum(nums []int, k int) int {
	res := 0
	hm := make(map[int]int)

	count := 0
	for _, val := range nums {
		count += val
		hm[count]++
		if k != 0 {
			res += hm[count-k]
		}
		if k == count {
			res += hm[count]
		}
	}

	return res
}

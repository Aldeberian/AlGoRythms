package leetcode

// You are given an integer array nums and an integer k.

// In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.

// Return the maximum number of operations you can perform on the array.

func maxOperations(nums []int, k int) int {
	res := 0
	count := make(map[int]int)

	for _, val := range nums {
		count[val]++
	}

	for _, val := range nums {
		if _, ok := count[k-val]; ok && count[k-val] > 0 {
			count[k-val]--
			res++
		}
	}

	return res / 2
}

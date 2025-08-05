package leetcode

// Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.

func uniqueOccurrences(arr []int) bool {
	store := make(map[int]int)

	for _, val := range arr {
		if _, ok := store[val]; ok {
			store[val]++
		} else {
			store[val] = 1
		}
	}

	other := make(map[int]bool)
	for _, val := range store {
		if other[val] {
			return false
		}
		other[val] = true
	}

	return true
}

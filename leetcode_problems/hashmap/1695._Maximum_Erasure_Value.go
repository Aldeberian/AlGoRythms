package leetcode

// You are given an array of positive integers nums and want to erase a subarray containing unique elements. The score you get by erasing the subarray is equal to the sum of its elements.

// Return the maximum score you can get by erasing exactly one subarray.

// An array b is called to be a subarray of a if it forms a contiguous subsequence of a, that is, if it is equal to a[l],a[l+1],...,a[r] for some (l,r).

func maximumUniqueSubarray(nums []int) int {
	subarray := make(map[int]bool)
	maxi := 0
	temp := 0
	start := 0

	for i, val := range nums {
		if subarray[val] {
			for j := start; j < i; j++ {
				start = j
				temp -= nums[j]
				subarray[nums[j]] = false
				if nums[j] == val {
					subarray[nums[j]] = true
					start++
					break
				}
			}
			temp += val
		} else {
			temp += val
			subarray[val] = true
		}
		maxi = max(temp, maxi)
	}

	return maxi
}

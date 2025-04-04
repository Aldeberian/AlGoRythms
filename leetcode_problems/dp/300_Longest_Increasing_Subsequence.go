package leetcode

// Given an integer array nums, return the length of the longest strictly increasing subsequence.

type Pair struct {
	Value int
	Count int
}

func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	tracking := []Pair{{nums[0], 1}}

	max := 1
	for _, val := range nums {
		tempMax := 1
		for _, val2 := range tracking {
			if val > val2.Value {
				if val2.Count >= tempMax {
					tempMax = val2.Count + 1

					if tempMax > max {
						max = tempMax
					}
				}
			}
		}
		tracking = append(tracking, Pair{Value: val, Count: tempMax})
	}
	return max
}

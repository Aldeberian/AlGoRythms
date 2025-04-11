package leetcode

import "sort"

// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	for i := range intervals {
		if i == len(intervals)-1 {
			return intervals
		}

		if intervals[i][1] >= intervals[i+1][1] {
			return merge(append(intervals[:i+1], intervals[i+2:]...))
		}

		if intervals[i][1] < intervals[i+1][1] && intervals[i][1] >= intervals[i+1][0] {
			intervals[i][1] = intervals[i+1][1]
			return merge(append(intervals[:i+1], intervals[i+2:]...))
		}
	}

	return [][]int{}
}

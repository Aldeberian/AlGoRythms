package leetcode

import (
	"sort"
)

// You are given a string s. We want to partition the string into as many parts as possible so that each letter appears in at most one part. For example, the string "ababcc" can be partitioned into ["abab", "cc"], but partitions such as ["aba", "bcc"] or ["ab", "ab", "cc"] are invalid.

// Note that the partition is done so that after concatenating all the parts in order, the resultant string should be s.

// Return a list of integers representing the size of these parts.

func partitionLabels(s string) []int {
	intervals := [][]int{}
	indices := make(map[rune]int)

	for i, val := range s {
		if indices[val] == 0 {
			indices[val] = len(intervals) + 1
			interval := []int{i, i}
			intervals = append(intervals, interval)
		} else {
			intervals[indices[val]-1][1] = i
		}
	}

	intervals = merge(intervals)

	res := []int{}
	for _, val := range intervals {
		res = append(res, val[1]-val[0]+1)
	}

	return res
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergedIntervals := make([][]int, 0, len(intervals))
	mergedIntervals = append(mergedIntervals, intervals[0])

	for _, interval := range intervals[1:] {
		if top := mergedIntervals[len(mergedIntervals)-1]; interval[0] > top[1] {
			mergedIntervals = append(mergedIntervals, interval)
		} else if interval[1] > top[1] {
			top[1] = interval[1]
		}
	}

	return mergedIntervals
}

package leetcode

import "math"

// Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.

// Return the smallest level x such that the sum of all the values of nodes at level x is maximal.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	count := make(map[int]int)
	res := 1
	maxi := math.MinInt
	maxD := 0

	var rec func(root *TreeNode, level int)

	rec = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		maxD = max(level, maxD)

		if _, ok := count[level]; ok {
			count[level] += root.Val
		} else {
			count[level] = root.Val
		}

		rec(root.Left, level+1)
		rec(root.Right, level+1)
	}

	rec(root, 1)

	for i := 1; i <= maxD; i++ {
		if count[i] > maxi {
			maxi = count[i]
			res = i
		}
	}

	return res
}

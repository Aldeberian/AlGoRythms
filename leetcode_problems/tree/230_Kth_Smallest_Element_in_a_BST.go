package leetcode

import "slices"

// Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.

func kthSmallest(root *TreeNode, k int) int {
	values := []int{}

	var rec func(node *TreeNode)

	rec = func(node *TreeNode) {
		if node == nil {
			return
		}

		values = append(values, node.Val)

		rec(node.Left)
		rec(node.Right)
	}
	rec(root)
	slices.Sort(values)

	return values[len(values)-k]
}

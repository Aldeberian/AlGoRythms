package leetcode

import "math"

// Given the root of a binary tree, determine if it is a valid binary search tree (BST).
// A valid BST is defined as follows:
//     The left subtree of a node contains only nodes with keys less than the node's key.
//
// 	   The right subtree of a node contains only nodes with keys greater than the node's key.
//
//     Both the left and right subtrees must also be binary search trees.

func isValidBST(root *TreeNode) bool {
	return recValid(root, math.MinInt64, math.MaxInt64)
}

func recValid(root *TreeNode, mini, maxi int) bool {
	if root == nil {
		return true
	}

	if root.Val >= maxi || root.Val <= mini {
		return false
	}

	nouvMin := max(mini, root.Val)
	nouvMax := min(maxi, root.Val)

	left := recValid(root.Left, mini, nouvMax)
	right := recValid(root.Right, nouvMin, maxi)

	return left && right
}

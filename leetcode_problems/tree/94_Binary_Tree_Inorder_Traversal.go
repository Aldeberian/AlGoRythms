package leetcode

// Given the root of a binary tree, return the inorder traversal of its nodes' values.

// Input: root = [1,null,2,3]
// Output: [1,3,2]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	var rec func(node *TreeNode)

	rec = func(node *TreeNode) {
		if node == nil {
			return
		}

		rec(node.Left)
		res = append(res, node.Val)
		rec(node.Right)
	}
	rec(root)

	return res
}

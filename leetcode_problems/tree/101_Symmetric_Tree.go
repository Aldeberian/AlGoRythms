package leetcode

import "reflect"

// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

func isSymmetric(root *TreeNode) bool {
	left := []int{}
	right := []int{}

	var rec func(root *TreeNode, isLeft bool)

	rec = func(root *TreeNode, isLeft bool) {
		if root == nil {
			if isLeft {
				left = append(left, 101)
			} else {
				right = append(right, 101)
			}
			return
		}
		if isLeft {
			left = append(left, root.Val)
			rec(root.Left, isLeft)
			rec(root.Right, isLeft)
		} else {
			right = append(right, root.Val)
			rec(root.Right, isLeft)
			rec(root.Left, isLeft)
		}
	}

	rec(root.Left, true)
	rec(root.Right, false)

	return reflect.DeepEqual(left, right)
}

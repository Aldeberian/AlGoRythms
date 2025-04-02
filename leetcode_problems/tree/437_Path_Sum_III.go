package leetcode

// Given the root of a binary tree and an integer targetSum, return the number of paths where the sum of the values along the path equals targetSum.

// The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).

func isSum(array []int, obj int) int {
	total := 0
	response := 0

	for i := len(array) - 1; i >= 0; i-- {
		total += array[i]
		if total == obj {
			response++
		}
	}

	return response
}

func pathSum(root *TreeNode, targetSum int) int {
	response := 0
	var rec func(root *TreeNode, path []int)

	rec = func(root *TreeNode, path []int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		response += isSum(path, targetSum)

		rec(root.Left, path)
		rec(root.Right, path)
	}
	rec(root, []int{})

	return response
}

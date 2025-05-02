package leetcode

// Given the root of a binary tree, flatten the tree into a "linked list":

//     The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
//     The "linked list" should be in the same order as a pre-order traversal of the binary tree.

func flatten(root *TreeNode) {
	nodes := make(map[int]*TreeNode)
	min := 0
	max := 0

	var rec func(root *TreeNode)

	rec = func(root *TreeNode) {
		if root == nil {
			return
		}
		nodes[root.Val] = root

		max = Max(root.Val, max)
		min = Min(root.Val, min)

		rec(root.Left)
		rec(root.Right)
	}

	retour := &TreeNode{Right: nodes[min]}
	curr := retour.Right
	for i := min + 1; i <= max; i++ {
		curr.Right = nodes[i]
		curr = curr.Right
	}

	return retour.Right
}

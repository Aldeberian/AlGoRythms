package leetcode

// Given the root of a binary tree, flatten the tree into a "linked list":

//     The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
//     The "linked list" should be in the same order as a pre-order traversal of the binary tree.

func flatten(root *TreeNode) {
	nodes := []*TreeNode{}

	var rec func(root *TreeNode)

	rec = func(root *TreeNode) {
		if root == nil {
			return
		}

		nodes = append(nodes, root)

		rec(root.Left)
		rec(root.Right)
	}

	rec(root)

	curr := root
	for i := 1; i < len(nodes); i++ {
		curr.Right = nodes[i]
		curr.Left = nil
		curr = curr.Right
	}
}

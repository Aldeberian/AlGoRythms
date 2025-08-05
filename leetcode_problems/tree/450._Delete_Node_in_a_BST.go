package leetcode

// Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

// Basically, the deletion can be divided into two stages:

//     Search for a node to remove.
//     If the node is found, delete the node

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findLeft(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}

	return findLeft(root.Left)
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		remplacement := findLeft(root.Right)
		root.Val = remplacement.Val
		root.Right = deleteNode(root.Right, remplacement.Val)
	}

	return root
}

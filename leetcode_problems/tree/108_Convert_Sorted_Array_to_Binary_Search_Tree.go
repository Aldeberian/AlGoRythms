package leetcode

// Given an integer array nums where the elements are sorted in ascending order, convert it to a height balanced binary search tree.

func sortedArrayToBST(nums []int) *TreeNode {
	middle := len(nums)/2 - 1
	right := len(nums)/2 + 1

	root := &TreeNode{Val: nums[middle+1]}

	curr := root
	for right <= len(nums)-1 {
		curr.Right = &TreeNode{Val: nums[right]}
		curr = curr.Right
		right++
	}

	curr = root
	for middle >= 0 {
		curr.Left = &TreeNode{Val: nums[middle]}
		curr = curr.Left
		middle--
	}

	return root
}

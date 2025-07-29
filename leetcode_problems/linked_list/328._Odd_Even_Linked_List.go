package leetcode

// Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

// The first node is considered odd, and the second node is even, and so on.

// Note that the relative order inside both the even and odd groups should remain as it was in the input.

// You must solve the problem in O(1) extra space complexity and O(n) time complexity.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curr := head
	currRight := head.Next
	rightHead := currRight

	for currRight != nil && currRight.Next != nil {
		curr.Next = curr.Next.Next
		currRight.Next = currRight.Next.Next

		curr = curr.Next
		currRight = currRight.Next
	}

	curr.Next = rightHead

	return head
}

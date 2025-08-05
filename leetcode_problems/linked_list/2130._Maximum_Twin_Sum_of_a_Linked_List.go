package leetcode

// In a linked list of size n, where n is even, the ith node (0-indexed) of the linked list is known as the twin of the (n-1-i)th node, if 0 <= i <= (n / 2) - 1.

//     For example, if n = 4, then node 0 is the twin of node 3, and node 1 is the twin of node 2. These are the only nodes with twins for n = 4.

// The twin sum is defined as the sum of a node and its twin.

// Given the head of a linked list with even length, return the maximum twin sum of the linked list.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	storage := make(map[int]*ListNode)
	n := 0
	curr := head

	for curr != nil {
		storage[n] = curr
		n++
		curr = curr.Next
	}

	maxi := 0
	for i := 0; i < n/2; i++ {
		maxi = max(maxi, storage[i].Val+storage[n-1-i].Val)
	}

	return maxi
}

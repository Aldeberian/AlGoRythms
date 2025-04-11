package leetcode

// Given the head of a linked list, remove the nth node from the end of the list and return its head.

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr := head
	total := 0

	for curr != nil {
		total++
		curr = curr.Next
	}

	if total == n {
		return head.Next
	}

	curr = head
	count := 0
	for curr != nil {
		count++
		if count == total-n-1 {
			if n > 1 {
				curr.Next = curr.Next.Next
			} else {
				curr.Next = nil
			}

			return head
		}
		curr = curr.Next
	}

	return head
}

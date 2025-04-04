package leetcode

// Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	prev := dummyHead
	curr := head

	for curr != nil && curr.Next != nil {
		temp := curr.Next.Next
		first := curr
		second := curr.Next

		prev.Next = second
		second.Next = first
		first.Next = temp

		prev = first
		curr = temp
	}

	return dummyHead.Next
}

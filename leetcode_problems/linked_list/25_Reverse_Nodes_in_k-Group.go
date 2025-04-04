package leetcode

// Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	prev := dummyHead
	curr := head

	for {
		tablo := []*ListNode{}
		for i := 0; i < k; i++ {
			if curr == nil {
				return dummyHead.Next
			}
			tablo = append(tablo, curr)
			curr = curr.Next
		}

		for i := len(tablo) - 1; i >= 0; i-- {
			prev.Next = tablo[i]
			prev = prev.Next
		}
		prev.Next = curr
	}
}

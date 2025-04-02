package leetcode

import (
	utils "AlGoRythms/utils"
)

// Given the head of a singly linked list, return true if it is a or false otherwise.

func isPalindrome(head *ListNode) bool {
	count := 0

	curr := head
	for curr != nil {
		count++
		curr = curr.Next
	}

	curr = head
	c2 := 0
	stack := utils.Stack[int]{}
	for c2 != count/2 {
		stack.Push(curr.Val)
		curr = curr.Next
		c2++
	}

	if count%2 == 1 {
		curr = curr.Next
	}

	for curr != nil {
		val, verif := stack.Pop()
		if verif == false || curr.Val != val {
			return false
		}
		curr = curr.Next
	}

	return true
}

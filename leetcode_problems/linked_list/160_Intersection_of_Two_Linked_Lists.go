package leetcode

// Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect. If the two linked lists have no intersection at all, return null.

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	countA, countB := 0, 0

	currA := headA
	for currA != nil {
		countA++
		currA = currA.Next
	}

	currB := headB
	for currB != nil {
		countB++
		currB = currB.Next
	}

	if countA == 0 || countB == 0 {
		return nil
	}

	currA = headA
	currB = headB

	for countA != countB {
		if countA > countB {
			currA = currA.Next
			countA--
		} else {
			currB = currB.Next
			countB--
		}
	}

	for currA != currB {
		currA = currA.Next
		currB = currB.Next

		if currA == nil || currB == nil {
			return nil
		}
	}

	return currB
}

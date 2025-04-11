package leetcode

// A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

// Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value
// of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the
// pointers in the original list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.

func copyRandomList(head *Node) *Node {
	store := make(map[*Node]*Node)

	curr := head
	for curr != nil {
		store[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		store[curr].Next = store[curr.Next]
		store[curr].Random = store[curr.Random]
		curr = curr.Next
	}

	return store[head]
}

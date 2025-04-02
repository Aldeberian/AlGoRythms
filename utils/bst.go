package utils

type BSTNode[T Comparable[T]] struct {
	Val         T
	Left, Right *BSTNode[T]
}

func (node *BSTNode[T]) Insert(val T) {
	compare := val.Compare(node.Val)

	if compare > 0 {
		if node.Right == nil {
			node.Right = &BSTNode[T]{Val: val}
		} else {
			node.Right.Insert(val)
		}
	} else if compare < 0 {
		if node.Left == nil {
			node.Left = &BSTNode[T]{Val: val}
		} else {
			node.Left.Insert(val)
		}
	}
}

func (root *BSTNode[T]) Search(val T) bool {
	if root == nil {
		return false
	}

	if val.Compare(root.Val) == 0 {
		return true
	}

	if val.Compare(root.Val) < 0 {
		return root.Left.Search(val)
	}
	return root.Right.Search(val)
}

func (root *BSTNode[T]) GetDepth() int {
	if root == nil {
		return 0
	}

	left := root.Left.GetDepth()
	right := root.Right.GetDepth()

	return max(left, right) + 1
}

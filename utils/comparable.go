package utils

type Comparable[T any] interface {
	// Compare must return
	// -1 if this < other
	//  0 if this == other
	//  1 if this > other
	Compare(other T) int
}

type Int int

func (i Int) Compare(other Int) int {
	if i < other {
		return -1
	} else if i > other {
		return 1
	}
	return 0
}

type Float float64

func (f Float) Compare(other Float) int {
	if f < other {
		return -1
	} else if f > other {
		return 1
	}
	return 0
}

type String string

func (s String) Compare(other String) int {
	if s < other {
		return -1
	} else if s > other {
		return 1
	}
	return 0
}

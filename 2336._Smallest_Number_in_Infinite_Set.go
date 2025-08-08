package leetcode

import "container/heap"

// You have a set which contains all positive integers [1, 2, 3, 4, 5, ...].

// Implement the SmallestInfiniteSet class:

//     SmallestInfiniteSet() Initializes the SmallestInfiniteSet object to contain all positive integers.
//     int popSmallest() Removes and returns the smallest integer contained in the infinite set.
//     void addBack(int num) Adds a positive integer num back into the infinite set, if it is not already in the infinite set.

type SmallestInfiniteSet struct {
	smallest int
	added    *IntHeap
	inHeap   map[int]bool
}

func Constructor() SmallestInfiniteSet {
	h := &IntHeap{}
	heap.Init(h)
	return SmallestInfiniteSet{0, h, map[int]bool{}}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.added.Len() > 0 {
		pop := heap.Pop(this.added).(int)
		this.inHeap[pop] = false
		return pop
	}

	this.smallest++
	return this.smallest
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num <= this.smallest && !this.inHeap[num] {
		this.inHeap[num] = true
		heap.Push(this.added, num)
	}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

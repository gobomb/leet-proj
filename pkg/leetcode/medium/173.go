package medium

import "container/heap"

/*
	173. Binary Search Tree Iterator
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	pq    IntHeap
	count int
}

func ConstructorBSTIterator(root *TreeNode) BSTIterator {
	rs, pq := []int{}, IntHeap{}

	preorderTraversal2(root, &rs)

	for _, r := range rs {
		heap.Push(&pq, r)
	}

	return BSTIterator{
		pq:    pq,
		count: len(rs),
	}
}

func preorderTraversal2(root *TreeNode, rs *[]int) {
	if root == nil {
		return
	}

	*rs = append(*rs, root.Val)

	preorderTraversal2(root.Left, rs)
	preorderTraversal2(root.Right, rs)
}

func (bsti *BSTIterator) Next() int {
	rs := heap.Pop(&bsti.pq).(int)
	bsti.count--

	return rs
}

func (bsti *BSTIterator) HasNext() bool {
	return bsti.count != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

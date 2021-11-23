package medium

import "log"

/*
	138. Copy List with Random Pointer
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

type randomNodeMaker struct {
	list   [][]int
	set    []*RandomNode
	length int
}

func makeRandomNode(list [][]int) *RandomNode {
	if len(list) == 0 {
		return nil
	}

	maker := randomNodeMaker{
		list:   list,
		set:    make([]*RandomNode, len(list)),
		length: len(list),
	}

	return maker.makeRandomNode(0)
}

func (r *randomNodeMaker) makeRandomNode(i int) *RandomNode {
	if i >= r.length || i < 0 {
		return nil
	}

	if r.set[i] != nil {
		return r.set[i]
	}

	if isNull(i) {
		return nil
	}

	node := &RandomNode{}
	r.set[i] = node

	node.Val = r.list[i][0]
	node.Next = r.makeRandomNode(i + 1)
	node.Random = r.makeRandomNode(r.list[i][1])

	return node
}

type randomCopyer struct {
	visited map[*RandomNode]*RandomNode
}

func (rc randomCopyer) copyRandomList(head *RandomNode) *RandomNode {
	if head == nil {
		return nil
	}

	if n, ok := rc.visited[head]; ok {
		return n
	}

	h := &RandomNode{}

	rc.visited[head] = h

	h.Val = head.Val
	if head.Next != nil {
		h.Next = rc.copyRandomList(head.Next)
	}

	if head.Random != nil {
		h.Random = rc.copyRandomList(head.Random)
	}

	return h
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Copy List with Random Pointer.
Memory Usage: 3.8 MB, less than 15.18% of Go online submissions for Copy List with Random Pointer.
*/

func copyRandomList(head *RandomNode) *RandomNode {
	return randomCopyer{
		visited: make(map[*RandomNode]*RandomNode),
	}.copyRandomList(head)
}

func (head *RandomNode) print() {
	h := head

	for h != nil {
		log.Println(h)
		h = h.Next
	}
}

// 剑指offer 26
/*
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Copy List with Random Pointer.
	Memory Usage: 3.5 MB, less than 100.00% of Go online submissions for Copy List with Random Pointer.
*/
func copyRandomList1(head *RandomNode) *RandomNode {
	if head == nil {
		return nil
	}
	h := head

	for h != nil {
		n := h.Next
		h.Next = &RandomNode{
			Val:  h.Val,
			Next: n,
		}
		h = n
	}

	h = head

	for h != nil {
		if h.Random != nil {
			h.Next.Random = h.Random.Next
		}

		h = h.Next.Next
	}

	h = head
	hn := head.Next
	nh := head.Next

	for nh != nil && h != nil {
		h.Next = nh.Next
		h = h.Next

		if nh.Next != nil {
			nh.Next = nh.Next.Next
		}

		nh = nh.Next
	}

	return hn
}

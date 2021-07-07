package leetcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	rs := []*Node{}
	levelReverseOrderTravel(root, 0, &rs)

	return root
}

func levelReverseOrderTravel(root *Node, depth int, rs *[]*Node) {
	if root == nil {
		return
	}
	if len(*rs) == depth {
		*rs = append(*rs, nil)
	}

	levelReverseOrderTravel(root.Right, depth+1, rs)

	root.Next = (*rs)[depth]
	(*rs)[depth] = root

	levelReverseOrderTravel(root.Left, depth+1, rs)
}

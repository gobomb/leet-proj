package leetcode

import (
	"log"
)

func levelOrder(root *TreeNode) [][]int {
	rs := [][]int{}
	getLevelOrder(root, 0, &rs)

	return rs
}

func getLevelOrder(root *TreeNode, depth int, rs *[][]int) {
	if root == nil {
		return
	}
	if len(*rs) == depth {
		*rs = append(*rs, []int{})
	}

	(*rs)[depth] = append((*rs)[depth], root.Val)

	getLevelOrder(root.Left, depth+1, rs)
	getLevelOrder(root.Right, depth+1, rs)
}

func levelOrderIter(node *TreeNode) [][]int {
	rs := [][]int{}

	queue := []*TreeNode{}

	if node != nil {
		queue = append(queue, node)
	}

	for i := 0; len(queue) != 0; i++ {
		rs = append(rs, []int{})

		log.Println(i, queue)

		p := []*TreeNode{}

		for len(queue) != 0 {
			r := queue[0]
			queue = queue[1:]

			if r == nil {
				continue
			}

			rs[i] = append(rs[i], r.Val)

			if r.Left != nil {
				p = append(p, r.Left)
			}
			if r.Right != nil {
				p = append(p, r.Right)
			}
		}

		queue = p
	}

	return rs
}

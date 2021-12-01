package medium

/*
	222. Count Complete Tree Nodes
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}

	for i := 0; i < len(queue) && queue[i] != nil; i++ {
		if queue[i].Left != nil {
			queue = append(queue, queue[i].Left)
		}

		if queue[i].Right != nil {
			queue = append(queue, queue[i].Right)
		}
	}

	return len(queue)
}

func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	cnt := 0

	for len(queue) != 0 {
		r := queue[0]
		cnt++
		queue = queue[1:]

		if r.Left != nil {
			queue = append(queue, r.Left)
		}

		if r.Right != nil {
			queue = append(queue, r.Right)
		}
	}

	return cnt
}

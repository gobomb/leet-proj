package medium

/*
	450. Delete Node in a BST
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	preRoot := &TreeNode{
		Right: root,
		Val:   100001,
	}
	deleteNode1(root, preRoot, false, key)

	return preRoot.Right
}

func deleteNode1(root, preRoot *TreeNode, lr bool, key int) {
	if root == nil {
		return
	}

	if root.Val == key {
		if root.Right != nil {
			if root.Right.Left != nil {
				alt, preAlt := findAlterRight(root.Right, root)
				root.Val = alt.Val
				preAlt.Left = alt.Right
			} else {
				root.Val = root.Right.Val
				root.Right = root.Right.Right
			}
		} else if root.Left != nil {
			if root.Left.Right != nil {
				alt, preAlt := findAlterLeft(root.Left, root)
				root.Val = alt.Val
				preAlt.Right = alt.Left
			} else {
				root.Val = root.Left.Val
				root.Left = root.Left.Left
			}
		} else {
			if lr {
				preRoot.Left = nil
			} else {
				preRoot.Right = nil
			}
		}

		return
	}

	deleteNode1(root.Left, root, true, key)
	deleteNode1(root.Right, root, false, key)
}

func findAlterLeft(root, preRoot *TreeNode) (*TreeNode, *TreeNode) {
	if root.Right != nil {
		return findAlterLeft(root.Right, root)
	}

	return root, preRoot
}

func findAlterRight(root, preRoot *TreeNode) (*TreeNode, *TreeNode) {
	if root.Left != nil {
		return findAlterRight(root.Left, root)
	}

	return root, preRoot
}

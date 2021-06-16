package leetcode

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if targetSum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}
	targetSum = targetSum - root.Val

	if hasPathSum(root.Left, targetSum) {
		return true
	}
	return hasPathSum(root.Right, targetSum)
}

/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func getChildrenSum(node *TreeNode, targetSum int, curSum int) bool {

	if node == nil {
		return false
	}

	if isLeafNode(node) && targetSum == curSum+node.Val {
		return true
	}

	leftResult := getChildrenSum(node.Left, targetSum, curSum+node.Val)
	rightResult := getChildrenSum(node.Right, targetSum, curSum+node.Val)

	return leftResult || rightResult
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return getChildrenSum(root, targetSum, 0)
}

// @lc code=end


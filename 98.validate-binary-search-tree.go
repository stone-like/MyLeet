/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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

func Traverse(node *TreeNode, min, max float64) bool {
	if node == nil {
		return true
	}

	v := float64(node.Val)

	return v < max && min < v && Traverse(node.Left, min, v) && Traverse(node.Right, v, max)
}

func isValidBST(node *TreeNode) bool {

	return Traverse(node, math.Inf(-1), math.Inf(1))
}

// @lc code=end


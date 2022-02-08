/*
 * @lc app=leetcode id=617 lang=golang
 *
 * [617] Merge Two Binary Trees
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
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	leftNewNode := mergeTrees(root1.Left, root2.Left)
	rightNewNode := mergeTrees(root1.Right, root2.Right)

	newTreeNode := &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  leftNewNode,
		Right: rightNewNode,
	}

	return newTreeNode
}

// @lc code=end


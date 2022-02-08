/*
 * @lc app=leetcode id=1448 lang=golang
 *
 * [1448] Count Good Nodes in Binary Tree
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

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func traverse(node *TreeNode, count *int, curMaxNum int) {
	if node == nil {
		return
	}

	if curMaxNum <= node.Val {
		*count++
	}

	traverse(node.Left, count, max(curMaxNum, node.Val))
	traverse(node.Right, count, max(curMaxNum, node.Val))
}

func goodNodes(root *TreeNode) int {
	count := 0

	traverse(root, &count, root.Val)

	return count
}

// @lc code=end


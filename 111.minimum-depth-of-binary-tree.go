/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMin := minDepth(root.Left)
	rightMin := minDepth(root.Right)

	getMinFn := func(left int, right int) int {

		if left == 0 {
			return right
		}

		if right == 0 {
			return left
		}

		min := left

		if right < left {
			min = rightMin
		}

		return min
	}

	//+1は自分の分
	return getMinFn(leftMin, rightMin) + 1
}

// @lc code=end


/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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

func Travarse(node *TreeNode, depthMap map[int][]int, depth int) {

	if node == nil {
		return
	}

	depthMap[depth] = append(depthMap[depth], node.Val)

	Travarse(node.Left, depthMap, depth+1)
	Travarse(node.Right, depthMap, depth+1)

}

func levelOrder(root *TreeNode) [][]int {

	node := root

	if node == nil {
		return nil
	}

	depthMap := make(map[int][]int)

	depth := 0

	Travarse(node, depthMap, depth)

	n := len(depthMap)

	ans := make([][]int, n)

	for k, v := range depthMap {
		ans[k] = v
	}

	return ans
}

// @lc code=end


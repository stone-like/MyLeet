/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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

func zigzagLevelOrder(root *TreeNode) [][]int {

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
		ans[k] = newSlice(v, k)
	}

	return ans
}

func newSlice(s []int, depth int) []int {
	if depth%2 == 0 {
		return s
	}

	return reverse(s)
}

func reverse(s []int) []int {
	temp := make([]int, 0, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		temp = append(temp, s[i])
	}

	return temp
}

// @lc code=end


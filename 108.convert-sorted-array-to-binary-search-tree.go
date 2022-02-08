/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
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

func makeNode(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	pivotIndex := len(nums) / 2
	leftIndex := pivotIndex - 1
	rightIndex := pivotIndex + 1

	indexMax := len(nums) - 1

	var leftNode *TreeNode
	var rightNode *TreeNode

	if rightIndex <= indexMax {
		rightNode = makeNode(nums[rightIndex:])
	}

	if 0 <= leftIndex {
		leftNode = makeNode(nums[:pivotIndex])
	}

	newNode := &TreeNode{
		Val: nums[pivotIndex],
	}

	newNode.Left = leftNode
	newNode.Right = rightNode
	return newNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return makeNode(nums)
}

// @lc code=end


/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	var evenNodes []*ListNode
	var oddNodes []*ListNode

	node := head
	index := 1
	for node != nil {
		if index%2 == 0 {
			evenNodes = append(evenNodes, node)
		} else {
			oddNodes = append(oddNodes, node)
		}

		index++
		node = node.Next
	}

	var ans *ListNode
	var ansHead *ListNode

	oddTurn := true

	for {
		if len(evenNodes) == 0 && len(oddNodes) == 0 {
			break
		}

		if oddTurn {
			ans.Next = 
		}

	}

	return ans.Next
}

// @lc code=end


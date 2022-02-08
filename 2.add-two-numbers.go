/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head

	canComputing := func(l1 *ListNode, l2 *ListNode) bool {
		if l1 != nil || l2 != nil {
			return true
		}

		//どっちのnodeもnilだったとしてもcarryがあればその分のNodeを作る必要があるのでcarryがあれば計算続行
		if carry != 0 {
			return true
		}

		return false
	}

	for canComputing(l1, l2) {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}

// @lc code=end


/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//ポイントはHeadとTailを作って二つのListNodeをくっつけられるようにすること
func partition(head *ListNode, x int) *ListNode {

	minNode := &ListNode{}
	maxNode := &ListNode{}
	minTail := minNode
	maxTail := maxNode
	cur := head
	for cur != nil {
		if cur.Val >= x {

			maxTail.Next = cur
			maxTail = maxTail.Next

		} else {

			minTail.Next = cur
			minTail = minTail.Next

		}

		cur = cur.Next
	}
	//最初のmaxNode,minNodeは空で、Nextに1->2->2みたいにつながっている
	minTail.Next = maxNode.Next
	maxTail.Next = nil

	return minNode.Next
}

// @lc code=end


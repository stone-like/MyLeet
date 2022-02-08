/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fastPointer := head
	slowPointer := head

	for {

		if fastPointer != nil && fastPointer.Next != nil {
			fastPointer = fastPointer.Next.Next
		} else {
			return false
		}

		//fastpointerはslowに先駆けてNodeの存在確認をしているのでslowでは存在確認はしなくてよい

		slowPointer = slowPointer.Next

		//比較したいのは構造体ではなくどのアドレスを指しているかなのでポインタのまま比較
		if fastPointer == slowPointer {
			return true
		}
	}

	return false
}

// @lc code=end


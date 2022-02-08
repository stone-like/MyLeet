/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	var behind *ListNode

	for head != nil {
		next := head.Next
		head.Next = behind

		behind = head
		head = next

	}

	return behind

}

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}

// 	cur := head

// 	listSlice := make([]*ListNode, 0)

// 	for {
// 		if cur == nil {
// 			break
// 		}

// 		listSlice = append(listSlice, cur)

// 		cur = cur.Next
// 	}

// 	for i := len(listSlice) - 1; i > 0; i-- {
// 		node := listSlice[i]
// 		prevNode := listSlice[i-1]

// 		node.Next = prevNode
// 	}
// 	//headに新しいreverseを当てはめる前にhead.Next = nil
// 	head.Next = nil

// 	head = listSlice[len(listSlice)-1]

// 	return head

// }

// @lc code=end


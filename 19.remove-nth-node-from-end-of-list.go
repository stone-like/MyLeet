/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//ポイントはダミーノードを使うこと、そうじゃないと{1}みたいな要素一つでこれを削除するとき辛すぎる
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummy := &ListNode{}
	dummy.Next = head
	left := dummy

	right := head
	for i := 1; i <= n; i++ {
		if right == nil {
			break
		}
		right = right.Next
	}

	for right != nil {

		right = right.Next
		left = left.Next
	}

	//dummyを使っているので{1}でもdummy->{1}でleft.Nextは{1}となるので、left.Nextがnilとなる心配はない

	left.Next = left.Next.Next

	return dummy.Next

}

// @lc code=end


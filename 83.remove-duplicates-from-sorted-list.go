/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//ソート済みなので一つ前の値だけ持っておけばいいはず
func Traverse(node *ListNode, prevVal int) *ListNode {
	if node == nil {
		return nil
	}

	if node.Val == prevVal {
		node := Traverse(node.Next, node.Val)
		return node
	}

	node.Next = Traverse(node.Next, node.Val)

	return node
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	//Valが-100<= val <=100なので存在しない値を最初は入れてあげる
	return Traverse(head, -101)
}

// @lc code=end


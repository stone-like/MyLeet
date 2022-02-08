/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	isCycle, slow := hasCycle(head)
	if !isCycle {
		return nil
	}
	//今回はCycle確認＋サイクルのスタートNodeの情報がいる
	//ただfastとslowが一致するのはcycleのどこかはわからないのでここでスタートに一致させる必要がある
	//こっちのやり方の方がメモリ的にはいいけどはっきり言って面倒
	//この解法が正しいことの証明
	//https://www.youtube.com/watch?v=Oz7-VlcTpSQ&ab_channel=%E3%82%B7%E3%83%AA%E3%82%B3%E3%83%B3%E3%83%90%E3%83%AC%E3%83%BC%E3%81%8B%E3%82%89%E3%80%81%E3%82%A8%E3%83%B3%E3%82%B8%E3%83%8B%E3%82%A2%E6%83%85%E5%A0%B1%E3%82%92%E5%B1%8A%E3%81%91%E3%81%BE%E3%81%99%21
	fast := head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func hasCycle(head *ListNode) (bool, *ListNode) {
	fast := head
	slow := head
	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true, slow
		}
	}
	return false, nil
}

// @lc code=end


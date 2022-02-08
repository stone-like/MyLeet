/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val > pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push define
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

// Pop define
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}

	pq := &PriorityQueue{}

	heap.Init(pq)
	//元のリストをそのまま入れると元のリストが弄られておかしくなるのでcopyを入れる、どうせNextは使わないし
	for _, l := range lists {
		node := l
		for {
			if node == nil {
				break
			}
			cloned := &ListNode{
				Val: node.Val,
			}

			heap.Push(pq, cloned)

			node = node.Next
		}
	}

	var prev *ListNode

	n := pq.Len()

	for i := 0; i < n; i++ {
		node := heap.Pop(pq).(*ListNode)
		if prev != nil {
			node.Next = prev
		}

		prev = node
	}

	return prev

}

// @lc code=end


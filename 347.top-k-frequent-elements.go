/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
// Item define
type Item struct {
	value int
	count int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push define
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

// Pop define
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}
	pq := &PriorityQueue{}
	heap.Init(pq)

	for k, v := range m {
		heap.Push(pq, &Item{
			value: k,
			count: v,
		})
	}

	result := make([]int, 0, k)

	for i := 0; i < k; i++ {
		item := heap.Pop(pq).(*Item)
		result = append(result, item.value)
	}

	return result

}

// @lc code=end


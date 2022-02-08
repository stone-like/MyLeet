/*
 * @lc app=leetcode id=621 lang=golang
 *
 * [621] Task Scheduler
 */

// @lc code=start
type PriorityQueue []*Task

type Task struct {
	count int
	time  int
}

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
	item := x.(*Task)
	*pq = append(*pq, item)
}

// Pop define
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

//countが大きいほどpriorityQueueを使って優先的に使うのがキモ
//使ったらHeap -> Queue
//Queue -> HeapはCoolDownになったら
func leastInterval(tasks []byte, n int) int {

	m := make(map[rune]int)

	for _, t := range tasks {
		m[rune(t)]++
	}
	//taskQueueをつくる
	maxHeap := &PriorityQueue{}

	heap.Init(maxHeap)

	var queue []*Task

	for _, count := range m {
		heap.Push(maxHeap, &Task{count: count, time: 0})
	}

	time := 0

	for {

		if maxHeap.Len() == 0 && len(queue) == 0 {
			break
		}

		time += 1

		if maxHeap.Len() != 0 {
			t := heap.Pop(maxHeap).(*Task)
			t.time = time + n
			t.count = t.count - 1

			if t.count != 0 {
				//こうすることでqueue -> heapにcount0のやつは絶対戻らない
				queue = append(queue, t)
			}

		}

		if len(queue) != 0 {
			if queue[0].time == time {
				heap.Push(maxHeap, queue[0])
				queue = queue[1:]
			}
		}

	}

	return time

}

// @lc code=end


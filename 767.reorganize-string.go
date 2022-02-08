/*
 * @lc app=leetcode id=767 lang=golang
 *
 * [767] Reorganize String
 */

// @lc code=start

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

//この問題の肝として、mapを作るのはそうなんだけど、そのままmapをrangeで回すと26^500にmaxでなってしまうのでだめ
//なのでmaxHeapをつかって取り出しを高速化する

func reorganizeString(s string) string {
	alphaMap := make(map[string]int)

	targetCount := 0
	for _, r := range s {
		target := string(r)

		alphaMap[target]++
		targetCount++

	}

	pq := make(PriorityQueue, len(alphaMap))

	i := 0

	for value, priority := range alphaMap {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}

	heap.Init(&pq)

	ret := ""
	var prev *Item

	for pq.Len() > 0 {

		target := heap.Pop(&pq).(*Item)

		ret += target.value
		target.priority--

		if prev != nil {
			heap.Push(&pq, prev)
			prev = nil
		}

		//一回heapからは出しっぱなしにして置いて、次のループの時また詰めなおす
		if target.priority != 0 {
			prev = target
		}

	}

	if len(ret) != targetCount {
		return ""
	}

	return ret

}

// @lc code=end


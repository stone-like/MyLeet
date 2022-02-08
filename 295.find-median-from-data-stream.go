/*
 * @lc app=leetcode id=295 lang=golang
 *
 * [295] Find Median from Data Stream
 */

// @lc code=start
//単純にソートだけだとO(N)でTLE,Heap一つだけだとソートできても取り出せない、なのでheapを二つ使いminHeap(LargeHeap),maxHeap(smallHeap)を使う
//minHeapの最小はmaxHeapの最大より常に大きいようにする
type LargeHeap []int

func (pq LargeHeap) Len() int {
	return len(pq)
}

func (pq LargeHeap) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq LargeHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push define
func (pq *LargeHeap) Push(x interface{}) {
	item := x.(int)
	*pq = append(*pq, item)
}

// Pop define
func (pq *LargeHeap) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

type SmallHeap []int

func (pq SmallHeap) Len() int {
	return len(pq)
}

func (pq SmallHeap) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq SmallHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push define
func (pq *SmallHeap) Push(x interface{}) {
	item := x.(int)
	*pq = append(*pq, item)
}

// Pop define
func (pq *SmallHeap) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

func abs(i int) int {
	if i > 0 {
		return i
	}

	return -1 * i
}

type MedianFinder struct {
	lh *LargeHeap
	sh *SmallHeap
}

func Constructor() MedianFinder {
	lh := &LargeHeap{}
	sh := &SmallHeap{}
	heap.Init(lh)
	heap.Init(sh)
	return MedianFinder{lh: lh, sh: sh}
}

func (this *MedianFinder) AddNum(num int) {
	//shに基本入れることにする
	//下記の条件の時は移動を行う
	heap.Push(this.sh, num)

	//Lenの差が1以上の時
	if abs(this.sh.Len()-this.lh.Len()) > 1 {
		curShMax := heap.Pop(this.sh).(int)
		heap.Push(this.lh, curShMax)
	}
	//shのmaxとlhのminが sh.max < lh.minとなっていないとき

	if this.sh.Len() != 0 && this.lh.Len() != 0 {
		shMax := (*this.sh)[0]
		lhMin := (*this.lh)[0]
		if shMax >= lhMin {
			curShMax := heap.Pop(this.sh).(int)
			heap.Push(this.lh, curShMax)
			curLhMin := heap.Pop(this.lh).(int)
			heap.Push(this.sh, curLhMin)
		}
	}

}

func (this *MedianFinder) FindMedian() float64 {
	//もしsh,lhのLenが等しくなければ、大きいLenを持つ方を使う

	shLen := this.sh.Len()
	lhLen := this.lh.Len()

	if shLen == 0 && lhLen == 0 {
		return 0
	}

	if shLen == 0 {
		return float64((*this.lh)[0])
	}

	if lhLen == 0 {
		return float64((*this.sh)[0])
	}

	curShMax := (*this.sh)[0]
	curLhMin := (*this.lh)[0]

	if shLen == lhLen {
		return float64(curShMax+curLhMin) / float64(2)
	}

	if shLen > lhLen {
		return float64(curShMax)
	}

	return float64(curLhMin)
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end


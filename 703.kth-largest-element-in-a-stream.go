//最小ヒープの問題、最小ヒープはIntSliceで実装されているのでこれを使う
//https://medium.com/@yasufumy/data-structure-heap-ecfd0989e5be
//goのheapについては下記
//https://note.com/kltl/n/nab029e7bcfe4

package _703_kth_largest_element_in_a_stream

/*  Heap, PriorityQueue */

import (
	"container/heap"
)

type KthLargest struct {
	k    int
	heap intHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := intHeap(nums)
	//Convert to heap
	heap.Init(&h)

	return KthLargest{
		k:    k,
		heap: h,
	}
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(&kl.heap, val)

	//Pop minimum element until len(h) = k and kthlargest = h[0]
	for len(kl.heap) > kl.k {
		//heapのpopで小さい順に取り除いていく,例えば[2,4,5,8]で3が追加された時、[2,3,4,5,8]で3番目は4だが、ここで2,3を取り除くと[4,5,8]となりheap[0]がk=3番目となる

		heap.Pop(&kl.heap)
	}

	return kl.heap[0]
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

/*
 * @lc app=leetcode id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */

// @lc code=start
//あらかじめ、二つのsliceがソートされていることを利用した時間対策版

//詳しい解説
//https://www.codingninjas.com/codestudio/problem-details/find-k-pairs-with-smallest-sums_1381413
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result, h := [][]int{}, &minHeap{}
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return result
	}
	if len(nums1)*len(nums2) < k {
		k = len(nums1) * len(nums2)
	}
	heap.Init(h)

	//まずnums1とnums[0]のsumだけをheapに入れる
	//これはnums1[i]をとった時に一番小さくなるのがnums2[0]だから
	for _, num := range nums1 {
		heap.Push(h, []int{num, nums2[0], 0})
	}
	for len(result) < k {
		min := heap.Pop(h).([]int)
		result = append(result, min[:2])
		//(nums1[i],nums2[0])からiが小さい順にとっていくのは良い
		//これは一番最初が(nums1[0],num2[0])なんだけど、このときnums1[0],nums2[0]が最小なのは間違いなくて、この次に大きいのは(nums1[0],nums2[1]) or (nums1[1],nums[0])となる
		//後者はもうheapに入っているので(nums1[0],nums2[1])をheapに追加

		//次のループでは(nums1[1],nums[0])か(nums1[0],nums2[1]) の小さい方がpopされる
		//この議論をk回続けるだけ
		//min[2]がなぜあるのかというと、
		//例えば,(nums1[0],nums2[1],1)がpopされたとする
		//この時次に小さいのは(nums1[0],nums2[2]) or (nums1[1],nums2[1])
		//ここで(nums1[0],nums2[2],2)が追加される
		//(nums1[1],nums2[1])はスルーで良いのか？と思うが、結局小さい順にとってきたいので良い
		//必要な時だけ新しいのを追加するということ

		//(nums1[1],nums2[1])が追加されるには (nums1[1],nums[0])がpopされなければならないが,
		//このペアがpopされないということは(nums1[0],nums2[2],2)の方が小さいということなのでpopされないならされないで構わない、必要な時にpop方式で問題ない
		//なので、nums2のindexだけを増やしていけばよくて、そのためmin[2]でnums2のindexを保持

		if min[2] < len(nums2)-1 {
			heap.Push(h, []int{min[0], nums2[min[2]+1], min[2] + 1})
		}
	}
	return result
}

type minHeap [][]int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i][0]+h[i][1] < h[j][0]+h[j][1] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Item define
// type Item struct {
// 	pair []int
// 	sum  int
// }

// // A PriorityQueue implements heap.Interface and holds Items.
// type PriorityQueue []*Item

// func (pq PriorityQueue) Len() int {
// 	return len(pq)
// }

// func (pq PriorityQueue) Less(i, j int) bool {
// 	return pq[i].sum > pq[j].sum
// }

// func (pq PriorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// }

// // Push define
// func (pq *PriorityQueue) Push(x interface{}) {
// 	item := x.(*Item)
// 	*pq = append(*pq, item)
// }

// // Pop define
// func (pq *PriorityQueue) Pop() interface{} {
// 	n := len(*pq)
// 	item := (*pq)[n-1]
// 	*pq = (*pq)[:n-1]
// 	return item
// }

// func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

// 	pq := &PriorityQueue{}

// 	heap.Init(pq)

// 	add := func(a, b int) {

// 		heap.Push(pq, &Item{
// 			pair: []int{a, b},
// 			sum:  a + b,
// 		})
// 		//そのままだとメモリ使い過ぎて、下記みたいになる、k個まで保有してればいい
// 		//runtime: out of memory: cannot allocate 71303168-byte block (469401600 in use)
// 		for pq.Len() > k {
// 			heap.Pop(pq)
// 		}
// 	}

// 	//time Exceedとなってしまうので、二重forはだめ
// 	for i := 0; i < len(nums1); i++ {
// 		for j := 0; j < len(nums2); j++ {
// 			add(nums1[i], nums2[j])
// 		}
// 	}

// 	n := k

// 	if pq.Len() < k {
// 		n = pq.Len()
// 	}

// 	//make([]int,0,n)としないのはappendせずにindexで値を入れるため
// 	//indexで値を入れているのは先頭に追加したいからで、appendで先頭追加は遅くなるため
// 	pair := make([][]int, n)

// 	for i := 0; i < n; i++ {
// 		item := heap.Pop(pq).(*Item)
// 		pair[n-1-i] = item.pair
// 	}

// 	return pair

// }

// @lc code=end


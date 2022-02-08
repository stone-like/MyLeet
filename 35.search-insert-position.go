/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
//二分探索の問題
//自作二分探索も一応書いた
// func abs(x int) int {
// 	return int(math.Abs(float64(x)))
// }

// func myBinarySearch(n int, f func(int) bool) int {
// 	upper := n
// 	lower := -1
// 	//二分探索は[lower...mid...upper]とあってmidをupperかlowerにすることで徐々に範囲を狭めていく
// 	//upper-lower <= 0になったら終わり
// 	for abs(upper-lower) > 1 {
// 		mid := (upper + lower) / 2

// 		if f(mid) {
// 			upper = mid
// 		} else {
// 			lower = mid
// 		}
// 	}

// 	return upper
// }

// func searchInsert(nums []int, target int) int {

// 	return myBinarySearch(len(nums), func(i int) bool {
// 		return nums[i] >= target
// 	})
// }

func searchInsert(nums []int, target int) int {
	sort.Ints(nums)
	index := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})

	return index
}

// @lc code=end


/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 */

// @lc code=start
//twopointer,binarySearchの問題、left,rightをそれぞれarrayの始めと終わりに置く
//arrayはsort済みなので、left+right > targetであれば
//小さくしたいのでrightを<-に
//left+right < targetであれば
//大きくしたいのでleftを->
func twoSum(numbers []int, target int) []int {

	if len(numbers) == 0 {
		return []int{}
	}
	//答えはあることが確定しているので == の時breakでよい
	left := 0
	right := len(numbers) - 1
	for {
		if numbers[left]+numbers[right] == target {
			//答えは1-indexなので+1してあげる
			return []int{left + 1, right + 1}
		}

		if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}

// @lc code=end


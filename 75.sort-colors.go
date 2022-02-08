/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
//別解、l,r,i三つのポインタを使う、lはLまでは全部0、rはRからは全部2を表す
//これで0の時はlを->に2のときはrを<-に,１のときはなにもしな
//iは2以外の時動く、なんで2のときはiを動かなさないかというと下記のエッジケースがあるから
//[0 1 2 1 0 2]
//   ↑ ↑   ↑
//   L i   R
//でiを動かすと
//[0 1 0 1 2 2]
//   ↑   ↑
//   L  i,R
//となるこうすると0とLが指している1のswapが行われない,なので
//[0 1 0 1 2 2]
//   ↑ ↑ ↑
//   L i R
//として同じiの位置でもう一回swapを試みることでうまくL方面のswapもうまくいく
func sortColors(nums []int) {

	l, r := 0, len(nums)-1
	i := 0
	for i <= r {
		if nums[i] == 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
		} else if nums[i] == 2 {
			nums[r], nums[i] = nums[i], nums[r]
			r--
			i-- //iを動かさないため
		}

		i++
	}
}

// func sortColors(nums []int) {

// 	for i := 1; i < len(nums); i++ {

// 		for nums[i] < nums[i-1] {
// 			j := i
// 			for j > 0 && nums[i] < nums[j-1] {
// 				j--
// 			}
// 			nums[i], nums[j] = nums[j], nums[i]
// 		}

// 	}
// }

// @lc code=end


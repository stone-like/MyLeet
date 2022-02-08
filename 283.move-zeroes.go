/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int) {

	if len(nums) == 0 {
		return
	}

	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}

			j++
		}
	}
	//jはleftMostZero
	//[1,3,0,0,4] index4までいってj=2
	//index 2の0 とindex4 の4をswapした時にもjを進める,lastZeroはindex3となる

	//0が連続しないパターン
	//[1,3,0,6,4] index4までいってこの時はすでに6と0がswapされているのですでにlastZeroはindex3となっている ↓
	//[1,3,6,0,4]

	//なので結局0が連続するパターンだけ気にすればいい

}

// @lc code=end


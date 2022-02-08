/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	res := nums[0]
	dp[0] = res

	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			//もし一個前のdpがマイナスだったら基準となるsubArrayをiからに更新
			dp[i] = nums[i]
		}

		res = max(res, dp[i])
	}

	return res

}

// @lc code=end


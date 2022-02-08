/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0], dp[1] = nums[0], max(nums[1], nums[0])
	//[2,7,9,3,1]として、dp[1]は[2,7]での問題に分割
	for i := 2; i < n; i++ {
		//ここでもiまでの問題に分割して考えている
		//dp[i-1]とnums[i]+dp[i-2]を比較しているのは例えば
		//i=3,value=3のときに
		//dp[2] = 11 ,nums[3]=3 + dp[1] = 7 = 10
		//で、nums[3]=3をとるなら直前の9が関与したdp[2]ではなく、9が関与しないdp[1]から考えなくてはいけない
		//小問題レベルで考えるといい
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[n-1]
}

// @lc code=end


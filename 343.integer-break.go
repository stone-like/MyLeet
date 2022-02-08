/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
func integerBreak(n int) int {
	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {

		for j := 1; j <= i-1; j++ {
			left := j
			right := i - j
			//ここでmaxをとっているのは3 = 2+1のときにdp[2]=1なので2を分解させないため
			//10 = 6+4 = 6 -> 3+3と分解させた時がいいときもあるのでmaxをとっている
			//leftもrightもnになることはないので、
			//n=4のとき、ans=4とそのまましてしまう心配もない
			leftValue := max(dp[left], left)
			rightValue := max(dp[right], right)
			dp[i] = max(dp[i], leftValue*rightValue)
		}
	}

	return dp[n]

}

// @lc code=end


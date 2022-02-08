/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1

	for i := 0; i <= n; i++ {
		//1step
		if i+1 <= n {
			dp[i+1] = dp[i+1] + dp[i]
		}

		//2step
		if i+2 <= n {
			dp[i+2] = dp[i+2] + dp[i]
		}
	}

	return dp[n]
}

// @lc code=end


/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func isOver(target, limit int) bool {
	return target > limit
}
func jump(nums []int) int {
	dp := make([]int, len(nums))

	minNum := len(nums)

	for i := 0; i < len(dp); i++ {
		dp[i] = minNum
	}

	dp[0] = 0

	for i, v := range nums {
		for j := 1; j <= v; j++ {
			if !isOver(i+j, len(nums)-1) {
				dp[i+j] = min(dp[i+j], dp[i]+1)
			}
		}
	}

	return dp[len(nums)-1]
}

// @lc code=end


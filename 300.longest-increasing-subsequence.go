/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

//これも後ろから見ていって小問題にしていく系
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for j := i; j < len(nums); j++ {
			if nums[i] < nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}

	ans := 1

	for _, v := range dp {
		ans = max(ans, v)
	}

	return ans
}

// @lc code=end


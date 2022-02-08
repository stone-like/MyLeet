/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {

	dp := make([]bool, len(nums))

	dp[0] = true
	for i := 0; i < len(nums); i++ {

		if dp[i] {
			for j := 1; j <= nums[i]; j++ {

				if i+j > len(nums)-1 {
					continue
				}
				dp[i+j] = true
			}
		}

	}

	return dp[len(nums)-1]
}

// @lc code=end


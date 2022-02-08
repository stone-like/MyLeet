/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */

// @lc code=start

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func houseRobber(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))

	dp[0] = nums[0]
	dp[1] = max(nums[1], dp[0]) //自分をとらない選択もある

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	//numsはcycleとなっているので、cycleをsub1,sub2に分ける
	sub1 := nums[:len(nums)-1]
	sub2 := nums[1:]

	return max(houseRobber(sub1), houseRobber(sub2))
}

// @lc code=end


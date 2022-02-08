/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minSubArrayLen(target int, nums []int) int {
	N := len(nums)
	left, right := 0, 0

	curSum := 0
	minLen, defaultMin := N*2, N*2

	for right <= N-1 {

		curSum += nums[right]

		for curSum >= target {
			minLen = min(minLen, right-left+1)
			curSum -= nums[left]
			left++
		}

		right++
	}

	if minLen == defaultMin {
		return 0
	}

	return minLen
}

// @lc code=end


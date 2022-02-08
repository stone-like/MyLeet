/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

// @lc code=start

//累積和を使う、動きとしては下記
// 例. [1,2,1,1,1]

// [0,1,3,4,5,6] 累積和の配列

// 後ろからdp[i]ごとにその前のやつを引いたやつを出していく

// dp[5] -dp[0] ,dp[1]... dp[4]
//dp[5] - dp[0]みたいに6 - 0みたいになるやつは最初からdp[i]まで全部足した奴
// dp[4] -dp[0] ,dp[1]... dp[3]
func subarraySum(nums []int, k int) int {

	if len(nums) == 0 {
		return 0
	}

	//まず累積和を出す
	dp := make([]int, 0, len(nums)+1)
	dp = append(dp, 0)

	for i, v := range nums {
		dp = append(dp, dp[i]+v)
	}

	count := 0

	for i := len(dp) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if dp[i]-dp[j] == k {
				count++
			}
		}
	}

	return count
}

// @lc code=end


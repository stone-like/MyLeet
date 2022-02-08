/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func min(a, b float64) float64 {
	if a < b {
		return a
	}

	return b
}

func coinChange(coins []int, amount int) int {

	dp := make([]float64, amount+1) //0~amountまで
	//dpのkeyは現時点でのamount,valueはamountを達成するのに必要なmin枚数

	for i := 0; i < len(dp); i++ {
		dp[i] = math.Inf(1)
	}

	dp[0] = 0

	for i := 0; i <= amount; i++ {
		for _, c := range coins {

			if i+c <= amount {
				dp[i+c] = min(dp[i+c], dp[i]+1)
			}
		}
	}

	if dp[amount] == math.Inf(1) {
		return -1
	}

	return int(dp[amount])

}

// @lc code=end


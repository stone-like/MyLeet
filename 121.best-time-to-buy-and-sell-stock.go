/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {

	n := len(prices)

	if n <= 1 {
		return 0
	}

	min := prices[0]
	maxProfit := 0

	for i := 1; i < n; i++ {
		if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}

		if prices[i] < min {
			//一番安い日に買いたいので
			min = prices[i]
		}
	}

	return maxProfit
}

// @lc code=end


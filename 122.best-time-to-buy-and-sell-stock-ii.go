/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
func maxProfit(prices []int) int {
	// [7,1,5,3,6,4]
	//上記で例えば1のときにsellして5の時にsellではなく,
	//1のときにsellして、6のときに売るみたいな方がいいんじゃないかと思うが,
	//これは[1,5,?,10]みたいになっていたとして、
	//2パターンあって、
	//1. [1,5,6,10] (10より小さい)
	//2. [1,5,11,10] (10より大きい)
	//2.の場合は議論を10 ->11に移せばいい、このときは1 ~5 , 5~11で買えばよい
	//1.の場合は10より小さいのでそのまま1~5,5~6,6~10で買えばいい
	//なので下記のような簡単な形でいい
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}

// @lc code=end


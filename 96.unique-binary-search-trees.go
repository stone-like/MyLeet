/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */

// @lc code=start
//ポイントは実際に構築するのではなくその時に構築できるBSTの個数はいくつか？を考えること
//n=1なら1,n=2なら2,
//n=3なら[1,2,3]で1,3が中央の時それぞれ2個、3が中央の時は一個
//これをn=4以降も続ければいい
//n=3までで分かるように、[1,2,3..]をnまでnが中央の時で考え、足していけばいい
//n=4だと
//1が中心 -> [2,3,4]が全部右にくる、左はなしで1パターン*5パターン
//2が中心 -> 1が左、[3,4]が右 1*2パターン
//3が中心 -> [1,2]が左、[4]が右 2*1パターン
//4が中心 -> [1,2,3]が全部左、5パターン
//全部足して14パターン
func numTrees(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1

	if n == 1 {
		return 1
	}

	for i := 2; i <= n; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			left := j
			right := i - j - 1
			sum += dp[left] * dp[right]
		}

		dp[i] = sum

	}

	return dp[n]

}

// @lc code=end


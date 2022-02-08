/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */

// @lc code=start

func isExpTwo(n int) bool {
	if n == 0 {
		return false
	}
	//２のべき乗の時は1000みたいに先頭のビットが一つだけ経つ
	//n-1をすると
	//例えば1000 ならn-1 = 0111となり&の結果は0000となる

	//例えば1100ならn-1 = 1011となり&の結果は1000となるので０にならない
	return (n & (n - 1)) == 0
}

func countBits(n int) []int {

	if n == 0 {
		return []int{0}
	}

	//2のべき乗でdpしていく
	// 1-> 2-> 4
	//            1の個数
	//0 - 0000  - 0
	//1 - 0001  - 1 + dp[n-1] //ここの-部分が2のべき乗の単位で変わる,nは0~nの数、n=1のときは1
	//2 - 0010  - 1 + dp[n-2]
	//3 - 0011  - 1 + dp[n-2]
	//4 - 0100  - 1 + dp[n-4]
	//5 - 0101  - 1 + dp[n-4]
	//6 - 0110  - 1 + dp[n-4]
	//7 - 0111  - 1 + dp[n-4]
	//8 -10000  - 1 + dp[n-8]
	dp := make([]int, n+1)
	expTwo := 1
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		if isExpTwo(i) {
			expTwo *= 2
		}
		dp[i] = 1 + dp[i-expTwo]
	}
	return dp
}

// @lc code=end


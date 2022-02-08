/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start

func uniquePaths(m int, n int) int {

	dp := make([][]int, m)

	//アイデアとして一番下の行から行単位で<-にpath数を求めていく
	//一番下と一番右は全部1で、m-2,n-2の部分だけを見ていく

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	//一番下の行は全部1で良い(一番下の行だけで考えれば→にしか行けないので)
	for i := 0; i < n; i++ {
		dp[m-1][i] = 1
	}
	//一番右の列も全部1
	for i := 0; i < m; i++ {
		dp[i][n-1] = 1
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
		}
	}

	return dp[0][0]
}

// @lc code=end


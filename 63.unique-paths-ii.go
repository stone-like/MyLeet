/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */

// @lc code=start

//今回は障害物があるのでゴールから求めるのではなくスタートから求めてDPしていく
//ゴールからやるやつだと到達不能な時が面倒
func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {

		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}

	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {

			dp[i][j] = dp[i-1][j] + dp[i][j-1]

			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			}
		}
	}
	return dp[m-1][n-1]

}

// @lc code=end


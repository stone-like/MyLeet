/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */

// @lc code=start
func numSquares(n int) int {

	dp := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		dp[i] = n + 1
	}

	dp[0] = 0
	//squareなものを順々に足していく、ただやってることわかりにくい
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			square := j * j
			//この短縮の仕方は前にもやったように結果と計算結果で逆算する的な奴
			if i-square < 0 {
				break
			}

			if 1+dp[i-square] < dp[i] {
				dp[i] = 1 + dp[i-square]
			}
		}
	}

	return dp[n]

}

//下記はTLE、単純に二重ループで良さそう？
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}

// 	return b
// }

// func dfs(dp []int, n int) int {

// 	if dp[n] != 100001 {
// 		return dp[n]
// 	}

// 	if isSquare(n) {
// 		dp[n] = 1
// 		return 1
// 	}

// 	for i := 1; i <= n-1; i++ {
// 		left := i
// 		right := n - i

// 		leftCount := dfs(dp, left)
// 		if dp[n] <= leftCount {
// 			continue
// 		}
// 		rightCount := dfs(dp, right)

// 		dp[n] = min(dp[n], leftCount+rightCount)
// 	}

// 	return dp[n]
// }

// func isSquare(num int) bool {
// 	target := math.Floor(math.Sqrt(float64(num)))

// 	return target*target == float64(num)
// }
// func numSquares(n int) int {

// 	dp := make([]int, n+1)

// 	for i := 0; i < n+1; i++ {
// 		dp[i] = 100001
// 	}

// 	dp[0] = 0
// 	dp[1] = 1

// 	return dfs(dp, n)

// }

// @lc code=end


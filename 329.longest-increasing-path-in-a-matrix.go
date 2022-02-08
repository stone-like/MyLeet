/*
 * @lc app=leetcode id=329 lang=golang
 *
 * [329] Longest Increasing Path in a Matrix
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func isBounded(y, x, boundY, boundX int) bool {
	return y <= boundY && 0 <= y && x <= boundX && 0 <= x
}

func isGreater(y, x int, prev int, matrix [][]int) bool {
	new := matrix[y][x]

	return prev < new
}

type Node struct {
	x int
	y int
}

func search(y, x int, prevVal int, dp map[Node]int, matrix [][]int) int {

	if !isBounded(y, x, len(matrix)-1, len((matrix)[0])-1) {
		return 0
	}

	if !isGreater(y, x, prevVal, matrix) {
		return 0
	}

	if length, exists := dp[Node{y, x}]; exists {
		return length
	}

	ret := 1

	for _, dir := range dirs {
		newY := y + dir[0]
		newX := x + dir[1]

		ret = max(ret, search(newY, newX, matrix[y][x], dp, matrix)+1)

	}

	dp[Node{y, x}] = ret

	return ret

}

func longestIncreasingPath(matrix [][]int) int {

	dp := make(map[Node]int)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			search(i, j, -1, dp, matrix)
		}
	}

	maxLen := -1

	for _, v := range dp {
		maxLen = max(maxLen, v)
	}

	return maxLen
}

// @lc code=end


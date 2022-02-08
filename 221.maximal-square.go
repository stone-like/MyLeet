/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 */

// @lc code=start

//一番右下から見ていくのがポイント
type RowCol struct {
	m int
	n int
}

func isOnBoard(y, x, m, n int) bool {
	return y >= 0 && m > y && x >= 0 && n > x
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func mins(nums ...int) int {

	tempMin := nums[0]
	for _, v := range nums {
		tempMin = min(tempMin, v)
	}

	return tempMin
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func dfs(y, x int, matrix [][]byte, cache map[RowCol]int) int {

	if !isOnBoard(y, x, len(matrix), len(matrix[0])) {
		return 0
	}

	key := RowCol{y, x}

	if v, exists := cache[key]; exists {
		return v
	}

	bottom := dfs(y+1, x, matrix, cache)
	diag := dfs(y+1, x+1, matrix, cache)
	right := dfs(y, x+1, matrix, cache)

	cache[key] = 0
	if string(matrix[y][x]) != "0" {
		cache[key] = 1 + mins(bottom, diag, right)
	}

	return cache[key]
}
func maximalSquare(matrix [][]byte) int {

	//key:座標、value:maxLengthofSquare
	cache := make(map[RowCol]int)

	dfs(0, 0, matrix, cache)

	ans := 0
	for _, v := range cache {
		ans = max(ans, v)
	}
	return ans * ans
}

// @lc code=end


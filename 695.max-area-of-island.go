/*
 * @lc app=leetcode id=695 lang=golang
 *
 * [695] Max Area of Island
 */

// @lc code=start

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

var dir = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)

	if m == 0 {
		return 0
	}

	n := len(grid[0])

	if n == 0 {
		return 0
	}
	//メモリ割り当てでエラーとなるので今回はvisitedを使う代わりに到達した奴は1->0にしてあげる

	maxNum := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res := searchIslands(grid, i, j)

				maxNum = max(maxNum, res)
			}
		}
	}

	return maxNum
}

func searchIslands(grid [][]int, y, x int) int {

	if !isOnBoard(grid, y, x) || grid[y][x] == 0 {
		return 0
	}

	grid[y][x] = 0 //1 -> 0へ

	count := 1

	for i := 0; i < 4; i++ {
		dx := x + dir[i][0]
		dy := y + dir[i][1]
		count += searchIslands(grid, dy, dx)
	}

	return count

}

func isOnBoard(grid [][]int, y, x int) bool {
	return y >= 0 && len(grid)-1 >= y && x >= 0 && len(grid[0])-1 >= x
}

// @lc code=end


/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func numIslands(grid [][]byte) int {

	//↓のlen
	m := len(grid)
	if m == 0 {
		return 0
	}
	//-> のlen
	n := len(grid[0])

	if n == 0 {
		return 0
	}

	res := 0
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	//byte を使うことについてだが、'1' -> 49となる
	//ただ'1'とはそのまま比較可能
	//runeの時 r -'0' や r - 'a'と合わせて覚えたい

	//ここでやっていることは未発見の1を探すこと、連結している１がsearchIslandsで見つからなくなったらres++してまた1探しに戻る
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				searchIslands(grid, &visited, i, j)
				res++
			}
		}
	}

	return res
}

func searchIslands(grid [][]byte, visited *[][]bool, y, x int) {
	(*visited)[y][x] = true

	for i := 0; i < 4; i++ {
		dx := x + dir[i][0]
		dy := y + dir[i][1]

		if isOnBoard(grid, dy, dx) && !(*visited)[dy][dx] && grid[dy][dx] == '1' {
			searchIslands(grid, visited, dy, dx)
		}
	}
}

func isOnBoard(grid [][]byte, y, x int) bool {
	return y >= 0 && len(grid)-1 >= y && x >= 0 && len(grid[0])-1 >= x
}

// @lc code=end


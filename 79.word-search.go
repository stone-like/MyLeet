/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
var dirs = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func isOnBoard(y, x int, board [][]byte) bool {
	return y >= 0 && len(board)-1 >= y && x >= 0 && len(board[0])-1 >= x
}

func dfs(y, x, currentIndex int, board [][]byte, target string, visited *[][]bool) bool {

	//len(word =5 で range dirs内でindex=4になっているときは0~4をたどって五文字すべて正しいときなので4+1=5の時にtrue
	if currentIndex == len(target) {
		return true
	}

	if !isOnBoard(y, x, board) || (*visited)[y][x] || target[currentIndex] != board[y][x] {
		return false
	}

	(*visited)[y][x] = true

	for _, dir := range dirs {

		dy := y + dir[0]
		dx := x + dir[1]

		ok := dfs(dy, dx, currentIndex+1, board, target, visited)

		if ok {
			return true
		}
	}
	//dfs内ではvisitedを保有したいが、この外側のm×nループではvisited=trueとしたままでいたくないので
	(*visited)[y][x] = false

	return false
}

func exist(board [][]byte, word string) bool {

	if len(board) == 0 {
		return false
	}
	m := len(board)
	n := len(board[0])

	visited := make([][]bool, m)

	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ok := dfs(i, j, 0, board, word, &visited)

			if ok {
				return true
			}
		}
	}

	return false
}

// @lc code=end


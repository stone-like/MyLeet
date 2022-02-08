/*
 * @lc app=leetcode id=909 lang=golang
 *
 * [909] Snakes and Ladders
 */

// @lc code=start
func IntToPos(x, y int, positionMap [][]int) int {
	return positionMap[y][x]
}

func PosToInt(pos int, positionMap [][]int) (y int, x int) {

	y = -1
	x = -1
	for i := 0; i < len(positionMap); i++ {
		for j := 0; j < len(positionMap[0]); j++ {
			if pos == positionMap[i][j] {
				return i, j
			}
		}
	}

	return
}

type Position struct {
	x    int
	y    int
	step int
}

func canJump(x, y int, board [][]int) bool {
	if board[y][x] != -1 {
		return true
	}

	return false
}

func bfs(target int, board [][]int, positionMap [][]int, queue []Position, visited map[int]bool) int {

	if len(queue) == 0 {
		return -1
	}
	current := queue[0]

	queue = queue[1:]
	curPos := IntToPos(current.x, current.y, positionMap)
	//もし遷移先がワープマスならワープ後のところにする
	if canJump(current.x, current.y, board) {
		curPos = board[current.y][current.x]
	}

	if curPos == target {
		return current.step
	}

	if visited[curPos] {
		return bfs(target, board, positionMap, queue, visited)
	}

	visited[curPos] = true

	for i := 1; i <= 6; i++ {
		newPos := curPos + i
		newY, newX := PosToInt(newPos, positionMap)
		if newX == -1 || newY == -1 {
			continue
		}

		queue = append(queue, Position{
			x:    newX,
			y:    newY,
			step: current.step + 1,
		})
	}

	return bfs(target, board, positionMap, queue, visited)

}

func reverse(target []int) []int {
	targetLen := len(target)
	temp := make([]int, targetLen)

	for i := 0; i < targetLen; i++ {
		temp[i] = target[targetLen-1-i]
	}

	return temp
}
func generatePositonMap(board [][]int) [][]int {

	positionMap := make([][]int, len(board))

	for i := 0; i < len(positionMap); i++ {
		positionMap[i] = make([]int, len(board[0]))
	}

	curNum := 1
	for i := len(board) - 1; i >= 0; i-- {
		for j := 0; j < len(board[0]); j++ {
			positionMap[i][j] = curNum
			curNum++
		}
	}

	//下から数えて偶数のところだけreverse

	for i := len(board) - 1; i >= 0; i-- {
		y := len(board) - 1 - i

		if y%2 != 0 {
			positionMap[i] = reverse(positionMap[i])
		}

	}

	return positionMap

}

func snakesAndLadders(board [][]int) int {
	var queue []Position
	visited := make(map[int]bool) //intはマス目
	positionMap := generatePositonMap(board)
	target := len(board) * len(board)
	queue = append(queue, Position{
		x:    0,
		y:    len(board[0]) - 1,
		step: 0,
	})
	return bfs(target, board, positionMap, queue, visited)
}

// @lc code=end


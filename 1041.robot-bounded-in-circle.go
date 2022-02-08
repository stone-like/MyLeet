/*
 * @lc app=leetcode id=1041 lang=golang
 *
 * [1041] Robot Bounded In Circle
 */

// @lc code=start

//この問題の肝は、positionが変わっていて、かつdirectionも変わっていればcycleができることを見抜けるかどうか、
//positionだけかわってもdirectionが変わらなければだめだし、
//directionだけ変わってもダメ
func isRobotBounded(instructions string) bool {
	dirX, dirY := 0, 1 //最初は北向き
	x, y := 0, 0

	for _, r := range instructions {
		if r == 'G' {
			x += dirX
			y += dirY
		} else if r == 'L' {
			dirX, dirY = -1*dirY, dirX
		} else {
			dirX, dirY = dirY, -1*dirX
		}
	}

	if x == 0 && y == 0 {
		return true
	}
	///ここに来るということは(x,y)!=(0,0)なのでpositionは変わってる
	//directionが変わっていればtrue
	if dirX != 0 || dirY != 1 {
		return true
	}
	//positionだけ変わっている1のはダメ
	return false
}

// @lc code=end


/*
 * @lc app=leetcode id=1871 lang=golang
 *
 * [1871] Jump Game VII
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func canReach(s string, minJump int, maxJump int) bool {
	var queue []int
	farthest := 0

	queue = append(queue, 0) //あらかじめ先頭のindexをqueueに追加

	for len(queue) != 0 {
		curIndex := queue[0]
		queue = queue[1:]
		//farthest以下はもうジャンプ済みなので、farthest+1から始めればよい
		//farthestがまだcurIndex+minJumpより小さかったらcurIndex+minJumpを使う
		start := max(curIndex+minJump, farthest+1)

		//windowはstartから今のindex+maxJump+1 || len(s)まで
		window := min(curIndex+maxJump+1, len(s))
		for i := start; i < window; i++ {
			if s[i] == '0' {
				queue = append(queue, i)
				if i == len(s)-1 {
					return true
				}
			}
		}

		farthest = curIndex + maxJump
	}

	return false
}

// @lc code=end


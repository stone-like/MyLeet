/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func abs(a int) int {
	if a < 0 {
		return -1 * a
	}

	return a
}

func isOk(mid, x int) bool {
	if mid*mid <= x {

		return true
	}

	return false
}
func mySqrt(x int) int {

	ok := 0
	ng := x

	if x == 1 {
		return 1
	}

	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		if isOk(mid, x) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

// @lc code=end


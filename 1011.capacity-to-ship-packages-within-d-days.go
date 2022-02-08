/*
 * @lc app=leetcode id=1011 lang=golang
 *
 * [1011] Capacity To Ship Packages Within D Days
 */

// @lc code=start
//二分探索を使うギリギリdaysに収まるようにする
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func shipWithinDays(weights []int, days int) int {

	ng := 0
	ok := 0

	for _, w := range weights {
		ok += w
	}

	isOK := func(capacity int) bool {

		count := 1
		current := 0

		for _, w := range weights {

			if w > capacity {
				return false
			}

			if current+w > capacity {
				count++
				current = w
				continue
			}

			current += w
		}

		return count <= days

	}

	for {
		if abs(ok-ng) <= 1 {
			break
		}

		mid := (ok + ng) / 2

		if isOK(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok

}

// @lc code=end


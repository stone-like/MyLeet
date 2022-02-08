/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
func abs(a int) int {
	if a < 0 {
		return -1 * a
	}

	return a
}
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	targetM := -1
	// targetN := -1

	left := 0
	right := m - 1
	//searchRow
	for abs(right-left) != 0 {
		mid := (right + left) / 2

		if matrix[mid][n-1] >= target {
			right = mid
		} else {
			left = mid + 1
		}

	}

	targetM = right

	left = 0
	right = n - 1
	for {

		mid := (right + left) / 2

		//条件判定する前に一回検査
		if matrix[targetM][mid] == target {
			// targetN = mid
			return true
		}

		if abs(right-left) == 0 {
			break
		}

		if matrix[targetM][mid] == target {
			// targetN = mid
			return true
		}

		if matrix[targetM][mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return false
}

// @lc code=end


/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start
func rotate(matrix [][]int) {
	l, r := 0, len(matrix)-1

	for {
		//4×4 -> 2×2みたいに外側から内側に動くfor
		if l >= r {
			break
		}
		//階層でのループ
		// [1,2,3]
		// [4,5,6]
		// [7,8,9]
		//だったら5以外の3×3が動く
		//ここのループ回数はN-1回、つまり3だったら2回
		//1×1はループなし
		//iは1 -> 2みたいに動かすときに使う

		for i := 0; i < r-l; i++ {
			top, bottom := l, r

			// topLeftを保存
			topLeft := matrix[top][l+i]

			// topLeft <- bottomLeft
			matrix[top][l+i] = matrix[bottom-i][l]

			// bottomLeft <- bottomRight
			matrix[bottom-i][l] = matrix[bottom][r-i]

			// bottomRight <- topRight
			matrix[bottom][r-i] = matrix[top+i][r]

			// topRight <- topLeft
			matrix[top+i][r] = topLeft
		}

		r--
		l++

	}
}

// @lc code=end


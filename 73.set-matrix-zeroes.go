/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

// @lc code=start
// timeはO(M*N)掛かってしまうが、SpaceをどうO(1)に落とすかがキモ
//O(1)解はO(M+N)の焼き直しみたいな感じで、
//O(M+N)解はGridのTopとLeftに余分に一行、一列作って、0がセットされる予定のところを0としてあげる、これでO(M+N)
//O(1)解はこの余分なtop,leftをそのままgridの[0][0~N](top),[0~M][0](left)に埋め込む
//ただtopの左端とleftの一番上がかぶるのでそこだけは作る必要がある、なのでO(1)
func setZeroes(matrix [][]int) {

	m := len(matrix)
	n := len(matrix[0])
	//0になってる場所は対応するTopとLeftを0にする
	//Leftの一番最初、Top側だけスペースがないので、LeftTopだけ作る
	leftTop := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {

				matrix[0][j] = 0

				if i == 0 {
					leftTop = 0
				} else {
					matrix[i][0] = 0
				}
			}

		}
	}

	//topLeftとleftTop以外
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	topLeft := matrix[0][0]

	if topLeft == 0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

	if leftTop == 0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	return
}

// @lc code=end


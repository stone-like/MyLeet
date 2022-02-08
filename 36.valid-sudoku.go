/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start

//ここでの肝はsquareMapをつくること3×3の小問題がsquareなんだけど、それぞれindex/3でどの小さいsquareにいるかを表すことができる
//例えばi=1,j=1だったら1/3=0,1/3=0で(0,0)の一番左上の3×3に存在することになる

type square struct {
	col int
	row int
}

func solve(board [][]byte) bool {

	cols := make(map[int]map[string]struct{})
	rows := make(map[int]map[string]struct{})
	squares := make(map[square]map[string]struct{})

	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			numStr := string(board[i][j])
			if numStr == "." {
				continue
			}

			square := square{
				col: i / 3,
				row: j / 3,
			}

			_, colMapExists := cols[i]
			_, rowMapExists := rows[j]
			_, squareMapExists := squares[square]

			if !colMapExists {
				cols[i] = make(map[string]struct{})
			}

			if !rowMapExists {
				rows[j] = make(map[string]struct{})
			}

			if !squareMapExists {
				squares[square] = make(map[string]struct{})
			}

			_, colExists := cols[i][numStr]
			_, rowExists := rows[j][numStr]
			_, squareExists := squares[square][numStr]

			if colExists || rowExists || squareExists {
				return false
			}

			cols[i][numStr] = struct{}{}
			rows[j][numStr] = struct{}{}
			squares[square][numStr] = struct{}{}
		}

	}

	return true
}

func isValidSudoku(board [][]byte) bool {

	return solve(board)
}

// @lc code=end


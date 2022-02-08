/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 */

// @lc code=start

//ポイントとしては現在の状況でQueenの支配下にあるマスをCol、posDiag,negDiagのSetで持つこと
//Queenは同じ行、同じ列に置けないので、一番上の行から一行ずつ下に下がってみていく
//posDiagは(Row+Column)が一定で表せて、
//negDiagは(Row-Column)が一定で表せる

func copy(s []string) []string {
	temp := make([]string, len(s))

	for i := 0; i < len(s); i++ {
		temp[i] = s[i]
	}

	return temp
}

func createRow(n int, str string) string {
	var s strings.Builder
	for i := 0; i < n; i++ {
		s.WriteString(str)
	}
	return s.String()
}

func canPlaceQueen(y, x int, rows, posDiag, negDiag map[int]struct{}) bool {
	_, alreadyAttackedInSameRow := rows[x]
	_, alreadyAttackedInPosDiag := posDiag[y+x]
	_, alreadyAttackedInNegDiag := negDiag[y-x]

	return !alreadyAttackedInSameRow && !alreadyAttackedInPosDiag && !alreadyAttackedInNegDiag
}

func dfs(ans *[][]string, current []string, curRow, queenNum int, rows, posDiag, negDiag map[int]struct{}) {
	if curRow == queenNum {
		*ans = append(*ans, copy(current))
		return
	}

	for i := 0; i < len(current); i++ {
		targetY := curRow
		targetX := i
		if !canPlaceQueen(targetY, targetX, rows, posDiag, negDiag) {
			continue
		}

		//Queenを追加するとき
		rows[targetX] = struct{}{}
		posDiag[targetY+targetX] = struct{}{}
		negDiag[targetY-targetX] = struct{}{}

		str := current[targetY]
		newStr := str[:targetX] + "Q" + str[targetX+1:]
		current[targetY] = newStr

		dfs(ans, current, curRow+1, queenNum, rows, posDiag, negDiag)

		//Queenを追加しないとき
		str = current[targetY]
		newStr = str[:targetX] + "." + str[targetX+1:]
		current[targetY] = newStr

		delete(rows, targetX)
		delete(posDiag, targetY+targetX)
		delete(negDiag, targetY-targetX)

	}
}

func solveNQueens(n int) [][]string {
	var ans [][]string

	current := make([]string, n)

	for i := 0; i < n; i++ {
		current[i] = createRow(n, ".")
	}

	rows := make(map[int]struct{})
	posDiag := make(map[int]struct{})
	negDiag := make(map[int]struct{})

	dfs(&ans, current, 0, n, rows, posDiag, negDiag)

	return ans
}

// @lc code=end


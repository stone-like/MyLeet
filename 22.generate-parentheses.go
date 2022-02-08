/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func dfs(res *[]string, stackStr string, lLimit, rLimit int, curStr string, endInd int) {

	if len(stackStr) != 0 && stackStr[len(stackStr)-1] == ')' {
		return
	}

	if len(curStr) == endInd && len(stackStr) == 0 {
		*res = append(*res, curStr)
		return
	}

	// "("を使う場合
	if lLimit != 0 {
		tempstackStr := stackStr + "("
		tempCurStr := curStr + "("
		//tempを使っているのは戻ってきたときに、下のrLimitに値が反映されてしまうから
		dfs(res, tempstackStr, lLimit-1, rLimit, tempCurStr, endInd)
	}

	// ")"を使う場合
	if rLimit != 0 && len(stackStr) > 0 {
		//こっちは変化が反映されるものがこれより下にないので問題ない
		//対となる"("を除去
		stackStr = stackStr[:len(stackStr)-1]
		curStr += ")"
		dfs(res, stackStr, lLimit, rLimit-1, curStr, endInd)
	}

}

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}

	var res []string
	stackStr := "("

	//lを一個使った状態から始める
	dfs(&res, stackStr, n-1, n, "(", n*2)

	return res
}

// @lc code=end


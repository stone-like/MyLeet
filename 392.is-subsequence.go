/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start

// func dfs(targetStr, sourceStr, curStr string, i int) bool {
// 	if targetStr == curStr {
// 		return true
// 	}

// 	if i >= len(sourceStr) {
// 		return false
// 	}
// 	//文字を使う場合
// 	tempStr := curStr + string(sourceStr[i])
// 	used := dfs(targetStr, sourceStr, tempStr, i+1)
// 	//文字を使わない場合
// 	unused := dfs(targetStr, sourceStr, curStr, i+1)
// 	return used || unused

// }
func isSubsequence(s string, t string) bool {

	if len(s) == 0 {
		return true
	}

	curSIndex := 0
	for i := 0; i < len(t); i++ {
		if s[curSIndex] == t[i] {
			curSIndex++
		}

		if curSIndex == len(s) {
			return true
		}

	}

	return false
}

// @lc code=end


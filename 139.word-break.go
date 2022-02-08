/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start

//これもLISとかと同じ後ろから見ていって小問題に分割するパターン
func wordBreak(s string, wordDict []string) bool {

	dp := make([]bool, len(s)+1)
	dp[len(s)] = true

	for i := len(s) - 1; i >= 0; i-- {
		for _, word := range wordDict {
			if i+len(word) <= len(s) && s[i:i+len(word)] == word {
				dp[i] = dp[i+len(word)]
				//例えばs=leetcode,wordDict=["leet","code"]とすると、
				//dpを一つ余分に8まで作り、0~8とし、dp[8]=trueとする、
				//wordごとに見ていくので,codeの部分までiは-されていく、
				//もし wordDictの中に一致するものがあれば,trueを伝播させていく、単にdp[i] = trueとするのではtrueが伝播せず小問題に分割できていないことになる

				//例　leetAAAAで今i=0としてleetがヒットしtrueになるとしても、
				//AAAA部分がヒットしていないのでだめ、なので後ろから見ていくにあたって、dpのtrueを伝播させる必要がある
			}

			if dp[i] == true {
				break
			}
		}
	}

	return dp[0]
}

// @lc code=end


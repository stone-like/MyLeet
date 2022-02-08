/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
//ブルートフォースでやるとO(N^3)掛かるので基本的なアイデアは
//回文の判定方法にある
//方法としては for でindexを -> に進めていくときに
//例えばcbabkでaを見るときだとして、aから両端をみてイコールなら回文というようにしていく
//これならO(N^2)
//エッジケースとしてcbbaみたいな偶数の時があるけどこのときは
//最初l=i,r=i+1から始める
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
func longestPalindrome(s string) string {

	res := ""
	resLen := 0

	for i := 0; i < len(s); i++ {
		//奇数
		l, r := i, i
		for {

			if l < 0 || r > len(s)-1 || s[l] != s[r] {
				break
			}

			tempRes := s[l : r+1]

			if resLen < len(tempRes) {
				res = tempRes
				resLen = len(tempRes)
			}

			l--
			r++

		}

		//偶数
		l, r = i, i+1
		for {

			if l < 0 || r > len(s)-1 || s[l] != s[r] {
				break
			}

			tempRes := s[l : r+1]

			if resLen < len(tempRes) {
				res = tempRes
				resLen = len(tempRes)
			}
			l--
			r++

		}
	}

	return res
}

// @lc code=end


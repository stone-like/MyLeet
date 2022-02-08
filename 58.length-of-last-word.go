/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	last := len(s) - 1
	first := last
	for first >= 0 && s[first] != ' ' {
		first--
	}

	return last - first
}

// @lc code=end


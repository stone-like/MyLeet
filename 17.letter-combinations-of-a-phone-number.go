/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func dfs(ans *[]string, digits string, curLetter string, m map[string][]string) {
	if len(digits) == 0 {
		*ans = append(*ans, curLetter)
		return
	}

	letters := m[string(digits[0])]

	for _, s := range letters {
		dfs(ans, digits[1:], curLetter+s, m)
	}

}
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var ans []string
	m := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	dfs(&ans, digits, "", m)
	return ans
}

// @lc code=end


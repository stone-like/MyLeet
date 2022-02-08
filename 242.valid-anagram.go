/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start

func isAnagram(s string, t string) bool {
	sr := []rune(s)
	tr := []rune(t)

	sort.Slice(sr, func(i, j int) bool {
		return sr[i] < sr[j]
	})
	sort.Slice(tr, func(i, j int) bool {
		return tr[i] < tr[j]
	})

	return string(sr) == string(tr)
}

// @lc code=end


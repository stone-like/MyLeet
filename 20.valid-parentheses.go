/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}

	correctMap := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := make([]rune, 0)

	for _, paren := range s {
		if paren == '(' || paren == '{' || paren == '[' {
			stack = append(stack, correctMap[paren])
			continue
		}

		//"){"みたいな形でstackに値が積まれる前にここにきてしまったらアウト
		if len(stack) == 0 {
			return false
		}

		latest := stack[len(stack)-1]

		if latest == paren {
			stack = stack[:len(stack)-1]
			continue
		}

		return false
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

// @lc code=end


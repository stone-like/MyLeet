/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

// @lc code=start
//二つの数字までしか機能しないというのがポイントで、
//121があるときに先頭から1 or 12を見ていくんだけどこれは1jump or 2jumpの選択肢と変換できる
//ただjumpの制約条件として1~26で無ければjumpできないという制約がある

func isValid(s string, m map[string]struct{}) bool {
	_, exists := m[s]
	return exists
}

func numDecodings(s string) int {

	m := make(map[string]struct{})

	for i := 1; i <= 26; i++ {
		ss := strconv.Itoa(i)
		m[ss] = struct{}{}
	}

	dp := make([]int, len(s)+1)

	if s[0] == '0' {
		return 0
	}

	dp[0] = 1
	for i := 0; i < len(s); i++ {

		//1jump
		single := string(s[i])

		if isValid(single, m) {
			dp[i+1] = dp[i+1] + dp[i]
		}

		//2jump
		if len(s)-1 > i {
			double := string(s[i]) + string(s[i+1])

			if isValid(double, m) {
				dp[i+2] = dp[i+2] + dp[i]
			}
		}

	}

	return dp[len(s)]
}

// @lc code=end


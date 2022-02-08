/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
func canJump(str string) bool {

	//001とか010とか01みたいな百桁に0があるときはだめ
	//唯一len(str) = 0のとき、str[0] = 0でもいい
	if len(str) != 1 && str[0] == '0' {
		return false
	}
	num, _ := strconv.Atoi(str)

	return 0 <= num && num <= 255

}

func dfs(str string, cur string, index int, ans *[]string, jumpNum int) {

	if jumpNum == -1 && len(cur) == len(str)+4 {
		*ans = append(*ans, strings.TrimLeft(cur, "."))
		return
	}

	if jumpNum == -1 {
		return
	}

	//jump1
	if index+1 <= len(str) && canJump(str[index:index+1]) {
		dfs(str, cur+"."+str[index:index+1], index+1, ans, jumpNum-1)
	}
	//jump2
	if index+2 <= len(str) && canJump(str[index:index+2]) {
		dfs(str, cur+"."+str[index:index+2], index+2, ans, jumpNum-1)

	}
	//jump3
	if index+3 <= len(str) && canJump(str[index:index+3]) {
		dfs(str, cur+"."+str[index:index+3], index+3, ans, jumpNum-1)
	}
}

func restoreIpAddresses(s string) []string {
	var ans []string
	if len(s) > 12 {
		return nil
	}
	dfs(s, "", 0, &ans, 3)
	return ans
}

// @lc code=end


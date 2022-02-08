/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
func firstUniqChar(s string) int {
	result := make([]int, 26)

	for _, v := range s {
		//v - 'a'でruneをintに変換し、- 'a'で26以内となる
		result[v-'a']++
	}

	for i, v := range s {
		if result[v-'a'] == 1 {
			return i
		}
	}

	return -1

}

// @lc code=end


/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)

	result := make([][]string, 0)

	for _, s := range strs {
		rs := []rune(s)
		sort.Sort(sortRunes(rs))

		strSlice := strMap[string(rs)]
		strSlice = append(strSlice, s) //nilに対してもappendできるので

		//新しいstrSliceで更新
		strMap[string(rs)] = strSlice

	}

	for _, v := range strMap {
		result = append(result, v)
	}

	return result

}

// @lc code=end


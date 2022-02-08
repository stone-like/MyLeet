/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})

	for _, v := range nums {
		m[v] = struct{}{}
	}

	ans := 0
	count := 0

	for _, v := range nums {
		target := v

		_, isLeftExists := m[v-1]

		if isLeftExists {
			continue
		}

		for {
			_, exists := m[target+1]
			if !exists {
				//最後にcount++(1->4で1からはじめて[1,2,3,4]がmapにあるとして、最後の4のカウント)
				count++
				ans = max(ans, count)
				break
			}

			count++
			target++

		}

		count = 0

	}

	return ans
}

// @lc code=end


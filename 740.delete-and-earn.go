/*
 * @lc app=leetcode id=740 lang=golang
 *
 * [740] Delete and Earn
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func generate(m map[int]int) []int {

	s := make([]int, len(m))

	index := 0
	for key := range m {
		s[index] = key
		index++
	}

	sort.Ints(s)

	return s

}

//ポイントはcountMapを作って、重複していなくてソート済みのsliceを作ること
//countMapの個数分一気に取り出してよく、->方向にやっていけば要素+1,要素-1のdeleteは気にしなくていい(遷移時に直前の要素が今の要素-1だったらmax(cur,prev)をとればいいので)
func deleteAndEarn(nums []int) int {
	countMap := make(map[int]int)

	for _, value := range nums {
		countMap[value]++
	}

	s := generate(countMap)
	dp := make([]int, len(s))

	dp[0] = countMap[s[0]] * s[0]

	for i := 1; i < len(s); i++ {

		val := s[i]
		earn := val * countMap[val]

		if s[i-1] != val-1 {
			dp[i] = dp[i-1] + earn
			continue
		}

		//この時前の要素と今の要素は共起できないので、
		//max(dp[i-1],newEarn+dp[i-2]で決める)

		if i-2 >= 0 {
			earn += dp[i-2]
		}
		dp[i] = max(dp[i-1], earn)
	}

	return dp[len(dp)-1]

}

// @lc code=end


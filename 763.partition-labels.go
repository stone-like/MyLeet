/*
 * @lc app=leetcode id=763 lang=golang
 *
 * [763] Partition Labels
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func partitionLabels(s string) []int {

	//mapのintはその文字が出たlastIndex,stringを進めていく内にこのlastIndex=endとなればendまでで重複がないということ
	//この方法はO(N)
	//stringが絡んでいく奴はあらかじめ何らかの形で答えにつながるmapを作るケースが多い、今回はlastIndexのmap。別のやつは*otみたいなmap
	m := make(map[string]int)
	for i, r := range s {
		m[string(r)] = i
	}

	end := 0
	size := 0
	var ans []int

	for i, r := range s {
		end = max(end, m[string(r)])
		size++

		if end == i {
			ans = append(ans, size)
			size = 0
		}
	}

	return ans
}

// func createMap(s string) map[string]struct{} {

// 	m := make(map[string]struct{})

// 	for _, r := range s {
// 		_, exists := m[string(r)]

// 		if exists {
// 			continue
// 		}

// 		m[string(r)] = struct{}{}
// 	}

// 	return m
// }

// func isDup(left, right map[string]struct{}) bool {
// 	for s, _ := range right {
// 		_, exists := left[s]

// 		if exists {
// 			return true
// 		}
// 	}

// 	return false
// }
// func partitionLabels(s string) []int {

// 	var ans []int
// 	//mapは一つでよくて、追加だけでいい(必ずrightにはleftにあるものが入っていないようにしたいから)
// 	end := len(s)
// 	for i := len(s) - 1; i >= 0; i-- {
// 		//後ろから見ていって右と左に分けて、右にだけあって左になくなるまで<-に動かす

// 		leftMap := createMap(s[i:end])
// 		rightMap := createMap(s[:i])

// 		if !isDup(leftMap, rightMap) {
// 			ans = append([]int{end - i}, ans...)
// 			end = i
// 		}

// 	}

// 	return ans
// }

// @lc code=end


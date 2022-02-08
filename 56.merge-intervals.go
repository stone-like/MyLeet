/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
func merge(intervals [][]int) [][]int {

	if len(intervals) == 0 {
		return nil
	}
	var ans [][]int

	//まずstart順にソート
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prevStart := intervals[0][0]
	prevEnd := intervals[0][1]

	for _, eachInterval := range intervals {
		start := eachInterval[0]
		end := eachInterval[1]

		if start <= prevStart && prevStart <= end && end <= prevEnd {
			//each [1,5] newInterval [2,7]のとき
			prevStart = min(prevStart, start)
			continue
		}

		if prevStart <= start && start <= prevEnd && prevEnd <= end {
			//each [6,10] newInterval [4,8]のとき
			prevEnd = max(prevEnd, end)
			continue
		}

		if prevStart <= start && end <= prevEnd {
			//すっぽり包む場合1
			//なにもしない
			prevStart = min(prevStart, prevStart)
			prevEnd = max(prevEnd, prevEnd)
			continue
		}

		if start <= prevStart && prevEnd <= end {
			//すっぽり包む場合2
			prevStart = min(prevStart, start)
			prevEnd = max(prevEnd, end)
			continue
		}

		//被らなくなったらそこのprevStart,prevEndを追加、
		//prevStart,prevEndを最新のものに
		ans = append(ans, []int{prevStart, prevEnd})
		prevStart = start
		prevEnd = end
		//追加するタイミングは [[1,3],[2,6],[8,10],[15,18]]の例としたら
		//[1,3][2,6]がそのままcontinueして[8,10]の時に[1,6]を追加
		//prevStart=8,prevEnd=10として次の[15,18]を追加
	}

	//ラストのprevStart,prevEndが被っていると追加されていないままなので最後に追加
	ans = append(ans, []int{prevStart, prevEnd})

	return ans
}

// @lc code=end


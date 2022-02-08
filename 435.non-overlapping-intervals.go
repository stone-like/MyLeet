/*
 * @lc app=leetcode id=435 lang=golang
 *
 * [435] Non-overlapping Intervals
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
func eraseOverlapIntervals(intervals [][]int) int {
	count := 0

	if len(intervals) == 0 {
		return count
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prevStart := intervals[0][0]
	prevEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]
		//今回 [1,2][2,3]みたいな形はoverlapしない

		//ソートしてるからこのケースはないと思うがprevStart,prevEnd = [2,3] start,end=[1,2]でoverlapしないように
		// if start <= prevStart && prevStart < end && end <= prevEnd {
		// 	//each [1,5] newInterval [2,7]のとき
		// 	prevStart = min(prevStart, start)
		// 	count++
		// 	continue
		// }
		//start,end = [2,3] prevStart,prevEnd=[1,2]でoverlapしないように
		//start < prevEndとしている
		if prevStart <= start && start < prevEnd && prevEnd <= end {
			//each [6,10] newInterval [4,8]のとき
			prevStart = min(prevStart, start)
			prevEnd = min(prevEnd, end)
			count++
			continue
		}

		if prevStart <= start && end <= prevEnd {
			//すっぽり包む場合1
			//なにもしない
			prevStart = min(prevStart, start)
			prevEnd = min(prevEnd, end)
			count++
			continue
		}

		if start <= prevStart && prevEnd <= end {
			//すっぽり包む場合
			prevStart = min(prevStart, start)
			prevEnd = min(prevEnd, end)
			count++
			continue
		}

		prevStart = start
		prevEnd = end
	}

	// for _, eachInterval := range intervals {
	// 	start := eachInterval[0]
	// 	end := eachInterval[1]

	// 	//今回[1,2] [2,3]のように片側がかぶっているものはoverlap判定しない

	// 	if start <= prevStart && prevStart <= end && end <= prevEnd {
	// 		//each [1,5] newInterval [2,7]のとき
	// 		prevStart = min(prevStart, start)
	// 		count++
	// 		continue
	// 	}

	// 	if prevStart <= start && start <= prevEnd && prevEnd <= end {
	// 		//each [6,10] newInterval [4,8]のとき
	// 		prevEnd = max(prevEnd, end)
	// 		count++
	// 		continue
	// 	}

	// 	// if prevStart <= start && end <= prevEnd {
	// 	// 	//すっぽり包む場合1
	// 	// 	//なにもしない
	// 	// 	prevStart = min(prevStart, prevStart)
	// 	// 	prevEnd = max(prevEnd, prevEnd)
	// 	// 	continue
	// 	// }

	// 	// if start <= prevStart && prevEnd <= end {
	// 	// 	//すっぽり包む場合2
	// 	// 	prevStart = min(prevStart, start)
	// 	// 	prevEnd = max(prevEnd, end)
	// 	// 	continue
	// 	// }
	// }

	return count
}

// @lc code=end


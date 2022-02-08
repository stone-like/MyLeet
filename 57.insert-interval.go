/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
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

func insert(intervals [][]int, newInterval []int) [][]int {

	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	var createdIntervals [][]int

	newStart := newInterval[0]
	newEnd := newInterval[1]

	curStart := newStart
	curEnd := newEnd

	//intervalが被っているとは次の条件の時、かぶっているときは即座にcreatedIntervalsに追加しない
	//1. A [1,5] B[2,7]
	//AStart <= BStart && AEnd <= BEnd
	//
	//2. A [2,7] B[1,5]
	//BStart <= AStart && BEnd <= AEnd
	//
	//3. A [6,7] B[4,8]
	//BStart <= AStart && AEnd <= BEnd
	for _, eachInterval := range intervals {
		start := eachInterval[0]
		end := eachInterval[1]

		if start <= newStart && newStart <= end && end <= newEnd {
			//each [1,5] newInterval [2,7]のとき
			curStart = min(curStart, start)
			continue
		}

		if newStart <= start && start <= newEnd && newEnd <= end {
			//each [6,10] newInterval [4,8]のとき
			curEnd = max(curEnd, end)
			continue
		}

		if newStart <= start && end <= newEnd {
			//すっぽり包む場合1
			curStart = min(curStart, newStart)
			curEnd = max(curEnd, newEnd)
			continue
		}

		if start <= newStart && newEnd <= end {
			//すっぽり包む場合2
			curStart = min(curStart, start)
			curEnd = max(curEnd, end)
			continue
		}

		createdIntervals = append(createdIntervals, eachInterval)

	}

	createdIntervals = append(createdIntervals, []int{curStart, curEnd})

	sort.Slice(createdIntervals, func(i, j int) bool {
		return createdIntervals[i][0] < createdIntervals[j][0]
	})

	return createdIntervals
}

// @lc code=end


/*
 * @lc app=leetcode id=853 lang=golang
 *
 * [853] Car Fleet
 */

// @lc code=start
//二つポインターを使う,endとstart
//positionと,speendをそーとして、
//ゴールに近いposition、つまり後ろから見ていく
//(goal - position) /speedでゴールまでの到達時間がわかり、
//start,endでもしEndGoalTime >= StartGoalTimeだったら衝突するのでfleetを形成する、その場合はstartを<-に進めてまた同じことをする

//EndGoalTime < StartGoalTimeだったらEndを<-にずらして同じようにする,このときansを++
type CarInfo struct {
	position int
	time     float64
}

func Calc(a, b int) float64 {
	return float64(a) / float64(b)
}

func carFleet(target int, position []int, speed []int) int {
	info := make([]CarInfo, len(position))
	for i := 0; i < len(position); i++ {
		info[i] = CarInfo{
			position: position[i],
			time:     Calc(target-position[i], speed[i]),
		}
	}
	sort.Slice(info, func(i, j int) bool {
		return info[i].position < info[j].position
	})

	ans := 1
	lastTime := info[len(info)-1].time
	for i := len(info) - 1; i >= 0; i-- {
		if info[i].time > lastTime {
			lastTime = info[i].time
			ans++
		}
	}

	return ans

}

// @lc code=end


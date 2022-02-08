/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start
//visitedMapを使うのだが肝はvisitedのremoveを使うこと
// 1 ->  3 -> 4
//       ->
//みたいに1 -> 4と 1->3 -> 4があって普通にvisitedを使うだけでは
//1 -> 4 のときにfalseになってしまう
//かといってvisitedを使わなければ  2 -> 0 -> 1 　　　　　　　　　　　　　　　　　　　　<-
//みたいなループが検知できなくなる
//なので自分のneighborを調べ終わった時に自身をremoveすればいい
//そうすれば 3 -> 4のとき4を調べ終わった時には4をremoveしていて、
// 1 -> 4のときでもOK
// 0 -> 1 -> 2 -> 0のときで0がvisitedにあるときに 2 -> 0では
//0がまだneighborを見切っていないので、 2->0のときはfalseとなる

func dfs(courseNum int, preMap map[int][]int, visited map[int]struct{}) bool {

	_, exists := visited[courseNum]
	if exists {
		return false
	}

	visited[courseNum] = struct{}{}

	for _, next := range preMap[courseNum] {

		ret := dfs(next, preMap, visited)

		if !ret {
			return false
		}
	}

	delete(visited, courseNum)
	delete(preMap, courseNum)

	return true
}
func canFinish(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int, numCourses)
	visited := make(map[int]struct{}, numCourses)

	for _, eachPre := range prerequisites {
		course := eachPre[0]
		preRequest := eachPre[1]
		preMap[preRequest] = append(preMap[preRequest], course)
	}

	for i := 0; i < numCourses; i++ {
		ret := dfs(i, preMap, visited)

		if !ret {
			return false
		}
	}

	return true
}

// @lc code=end


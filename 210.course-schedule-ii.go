/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 */

// @lc code=start

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

func dfs(visited, cycle map[int]struct{}, preMap map[int][]int, ans *[]int, curCourse int) bool {

	_, hasCycle := cycle[curCourse]
	if hasCycle {
		return false
	}

	_, visit := visited[curCourse]

	if visit {
		return true
	}

	//サイクルはdfsを再帰しているときにループ検知に使う、なので行きがけでなくてはダメ
	cycle[curCourse] = struct{}{}

	neighbor := preMap[curCourse]

	for _, v := range neighbor {
		ok := dfs(visited, cycle, preMap, ans, v)
		if !ok {
			return false
		}
	}

	//終わりがけにcycleを削除しないと、例えば
	// 0 -> 1 -> 3
	//   -> 2 ->
	//のような形で2->3行くときにfalseになってしまう
	delete(cycle, curCourse)

	//visitedは一番外のループのdfsを回して要る部分で何度も同じ場所見ないようにする、なので終わりがけでいい
	//cycleと違ってvisitedは終わり崖に追加するので削除しないでいい
	visited[curCourse] = struct{}{}
	*ans = append(*ans, curCourse)

	return true

}

func findOrder(numCourses int, prerequisites [][]int) []int {

	ans := make([]int, 0, numCourses)

	preMap := make(map[int][]int)
	visited := make(map[int]struct{})
	cycle := make(map[int]struct{})

	for i := 0; i < numCourses; i++ {
		preMap[i] = nil
	}
	for _, pair := range prerequisites {
		start := pair[1]
		end := pair[0]

		preMap[start] = append(preMap[start], end)
	}

	for k, _ := range preMap {
		ok := dfs(visited, cycle, preMap, &ans, k)
		if !ok {
			return nil
		}
	}

	reverse(ans)

	return ans

}

// @lc code=end


/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
//dfsを使ったbacktrackで解く
//下記が詳しい
// https://www.youtube.com/watch?v=REOH22Xwdkk&t=5s

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func Copy(s []int) []int {
	temp := make([]int, len(s))

	for i, v := range s {
		temp[i] = v
	}

	return temp
}

func dfs(i int, nums, subset []int, res *[][]int) {
	//DecisionTreeのLeafに到達した時はappend
	if i >= len(nums) {
		//copyしているのはせっかくresに追加してもsliceが変更されてしまったら台無しなため
		*res = append(*res, Copy(subset))
		return
	}

	//追加する
	subset = append(subset, nums[i])
	dfs(i+1, nums, subset, res)

	//追加しない
	subset = subset[:len(subset)-1]
	dfs(i+1, nums, subset, res)
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0, Pow(2, len(nums)))
	// dfs := func(i int) {

	// 	//DecisionTreeのLeafに到達した時はappend
	// 	if i >= len(nums) {
	// 		res = append(res, Copy(subset))
	// 		return
	// 	}

	// 	//追加する
	// 	subset = append(subset, nums[i])
	// 	dfs(i + 1)

	// 	//追加しない
	// 	subset = subset[:len(subset)-1]
	// 	dfs(i + 1)
	// }

	dfs(0, nums, nil, &res)

	return res
}

// @lc code=end


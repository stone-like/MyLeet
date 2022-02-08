/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
//[2,3,6,7]を例とするとブルートフォースで解こうとするとdecisionTreeの一つの択が4つになってしまう、そして重複を検知するのが難しい
//なのでbackTrackで解く,backTrackにすると配列の先頭から2を最低一つは含むときと一つも2を含まないときみたいに分けていく

//下記が詳しい
// https://www.youtube.com/watch?v=GBKI9VSKdGg
func Copy(s []int) []int {
	temp := make([]int, len(s))

	for i, v := range s {
		temp[i] = v
	}

	return temp
}

func dfs(candidates []int, target int, currentIndex int, total int, sumSlice []int, res *[][]int) {

	if total == target {
		*res = append(*res, Copy(sumSlice))
		return
	}

	if currentIndex > len(candidates)-1 || total > target {
		return
	}

	//制限無しの時,制限ないのでcurrentIndexは変化なし
	sumSlice = append(sumSlice, candidates[currentIndex])
	dfs(candidates, target, currentIndex, total+candidates[currentIndex], sumSlice, res)

	//candidateからcurrentIndexを使わない方へ分岐(この分岐では絶対curIndexのやつは使えない)
	sumSlice = sumSlice[:len(sumSlice)-1]
	dfs(candidates, target, currentIndex+1, total, sumSlice, res)
}
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	dfs(candidates, target, 0, 0, nil, &res)

	return res

}

// @lc code=end


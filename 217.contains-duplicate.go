/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {

	if len(nums) == 0 {
		return false
	}
	//ソートする(O(nlogn))
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	//ソート済みなのでもし重複があればnextと同じになる部分が出る
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}

	//もしくはmapでset作ってnumsとmapのlenを比べるか、ただmapを使うと1<=nums.length<=10^5なので結構メモリ食ってしまうのがネック

	return false

}

// @lc code=end


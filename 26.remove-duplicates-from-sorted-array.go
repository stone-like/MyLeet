/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {

	if len(nums) < 2 {
		return len(nums)
	}
	resultLen, preVal := 1, nums[0]

	//swapしていくだけ
	for i := 1; i < len(nums); i++ {
		if nums[i] != preVal {
			nums[resultLen] = nums[i]
			preVal = nums[i]
			resultLen++
		}
	}
	return resultLen

}

// @lc code=end


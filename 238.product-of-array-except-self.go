/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {

	if len(nums) == 0 {
		return nil
	}
	output := make([]int, len(nums))

	//insert prefix
	//nums[0]のprefixはないので1を入れてあげる
	output[0] = 1
	for i := 1; i < len(nums); i++ {
		output[i] = output[i-1] * nums[i-1]
	}

	//insert suffix and multiple
	//nums[len(nums)-1]のsuffixは無いので1かけてあげる、なので処理なし

	postSum := 1
	for i := len(nums) - 2; i >= 0; i-- {
		postSum *= nums[i+1]
		output[i] = output[i] * postSum
	}

	return output
}

// @lc code=end


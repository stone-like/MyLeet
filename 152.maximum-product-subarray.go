/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//連続したところをとるとあるが、最初の二つのindex内の小問題に分割して次のindexに移っていくという手法を取れる
//[2,3,-2...]
//とあって、[2,3]で最初min,max = 2
//min* nums[i],max * nums[i],nums[i]の中からmin,maxをとる
//すると6,6,3となり、
//max=6,min=3となる
//連続したという条件なので、かけたものを使うか、それともまるきり使わないか
//-2のときは
//-12,-6,-2となり
//max -2 ,min -12となる
func maxProduct(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	res := -11

	curMin, curMax := 1, 1

	for i := 0; i < len(nums); i++ {

		if nums[i] == 0 {
			curMax = 1
			curMin = 1
			res = max(0, res)
			continue
		}

		tempMax := max(nums[i]*curMax, max(nums[i]*curMin, nums[i]))
		tempMin := min(nums[i]*curMax, min(nums[i]*curMin, nums[i]))

		curMax = tempMax
		curMin = tempMin

		res = max(curMax, res)
	}

	return res
}

// @lc code=end


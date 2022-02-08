/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
// Boyer–Moore majority vote algorithm
//ただし、slice内にmajorityがなくてもreturnとしてsliceの値を返してしまうので注意、二分探索みたいな感じ
//例えば[2,2,1,1,3]のときとかは最後の3が答えとして出てしまうので、そもそもinputにmajorityが存在しない時は使えないという認識の方がよさそう
func majorityElement(nums []int) int {

	res, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if res == nums[i] {
			count++
			continue
		}

		if count == 0 {
			res = nums[i]
			count = 1
			continue
		}

		count--
	}

	return res
}

// @lc code=end


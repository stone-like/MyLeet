/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {

	cache := make(map[int]int) //value -> index

	for k, v := range nums {
		//任意の2つを足してtargetになる-> v + cacheValue = target
		// cacheにtarget - v となる値があればそのcacheの値とvを足せばtargetになるので即returnでいい
		targetInd, ok := cache[target-v]
		if ok {
			return []int{targetInd, k}
		}
		cache[v] = k
	}

	return []int{}

}

// func twoSum(nums []int, target int) []int {

// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			v := nums[i] + nums[j]
// 			if v == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}

// 	return []int{}

// }

// @lc code=end


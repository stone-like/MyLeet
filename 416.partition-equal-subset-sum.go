/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */

// @lc code=start
func canPartition(nums []int) bool {
	//後ろから見ていく
	//[1,5,11,5]だったらまず一番後ろの５から見ていって小問題に分解する
	//条件としてはsum(nums)/2 =targetに対して、
	//[1 ,(5,11,5)]で5,11,5までみたとして、
	//  t in subset == targteか t+1 == targetならtrue

	if len(nums) == 0 {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	dp := make(map[int]struct{})
	dp[0] = struct{}{} //最初のlen(nums)-1の時にdpがrangeで回るように,0はそのindexまで何もsumに使わなかったときを表す
	//dpにはそのindexまでのsubsetのsumをすべて入れている
	for i := len(nums) - 1; i >= 0; i-- {
		nextDp := make(map[int]struct{})
		for v, _ := range dp {
			nextDp[nums[i]+v] = struct{}{}
			nextDp[v] = struct{}{}
		}

		dp = nextDp
	}

	for v, _ := range dp {
		if v == target {
			return true
		}
	}

	return false

}

// @lc code=end


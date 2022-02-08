/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
func refbit(num uint32, pos uint) int {
	return int((num >> pos) & 1)
}

func hammingWeight(num uint32) int {
	ans := 0
	for i := 0; i < 32; i++ {
		bit := refbit(num, uint(i))

		if bit == 1 {
			ans++
		}
	}
	return ans
}

// @lc code=end


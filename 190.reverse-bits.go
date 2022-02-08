/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	res := 0
	for i := 0; i < 32; i++ {
		next := num >> i & 1
		res = res << 1
		res = int(uint32(res) + next)
	}
	return uint32(res)
}

// @lc code=end


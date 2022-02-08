/*
 * @lc app=leetcode id=371 lang=golang
 *
 * [371] Sum of Two Integers
 */

// @lc code=start
func getSum(a int, b int) int {
	xor := a ^ b
	carry := (a & b) << 1

	for {
		if carry == 0 {
			return xor
		}

		newxor := xor ^ carry
		newcarry := (xor & carry) << 1

		xor = newxor
		carry = newcarry
	}

	return xor
}

// @lc code=end


/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
func Traverse(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	res := Traverse(x, n/2)

	if n%2 == 0 {
		return res * res
	}

	//奇数の時はbaseが一つ足りない(x^5 = x^2 * x^2 * x)
	return res * res * x
}

func myPow(x float64, n int) float64 {
	res := Traverse(x, abs(n))

	if n <= 0 {
		return 1 / res
	}

	return res
}

// @lc code=end


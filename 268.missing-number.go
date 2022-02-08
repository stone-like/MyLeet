/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
//XOR Answer
func missingNumber(nums []int) int {
	//xorの特性として計算順序が関係ないこと、同じ数をxorするとゼロになることをつかう
	//[0,1,2,3] xor [0,1,2]だったらどんな計算順序でやってもよくて、3だけが消えないので3が答えということになる(0,1,2はxorで消える)

	xor := 0

	for _, v := range nums {
		xor ^= v
	}

	for i := 0; i <= len(nums); i++ {
		xor ^= i
	}

	return xor
}

//Sum Answer
// func missingNumber(nums []int) int {
// 	sum := 0

// 	for _, v := range nums {
// 		sum += v
// 	}

// 	expectSum := 0
// 	for i := 0; i <= len(nums); i++ {
// 		expectSum += i
// 	}

// 	return expectSum - sum
// }

// @lc code=end


/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */

// @lc code=start
//配列の順序を変えて次に大きい数になるようにする
//1. まず配列をprefixとsuffix部分にわける
//prefixは降順になっている部分、suffixは昇順になっている部分
//[6,2,1,5,4,3,0]
//だと6,2,1までがprefix、5,4,3,0がsuffix
//2. prefixの最後をkとして、
//nums[k] < nums[suffix]となる最小のindexをlとする
//kとlをswap
//[6,2,3,5,4,1,0]
//3. 最後にsuffix部分をreverseして終わり
//[6,2,3,0,1,4,5]

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

func nextPermutation(nums []int) {

	k := len(nums) - 2
	for {
		if k < 0 {
			break
		}

		if nums[k+1] > nums[k] {
			break
		}
		k--
	}

	if k == -1 {
		//reverse
		reverse(nums)
		return
	}

	for i := len(nums) - 1; i > k; i-- {
		if nums[i] > nums[k] {
			nums[k], nums[i] = nums[i], nums[k]
			break
		}
	}

	reverse(nums[k+1:])

	return

}

// @lc code=end


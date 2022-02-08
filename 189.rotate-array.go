/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func swap(start, end int, nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

func rotate(nums []int, k int) {
	//二重reverse
	//[1,2,3,4,5,6,7], k = 3だったら
	//まず一回目全体をreverse
	//[7,6,5,4,3,2,1]
	//二回目は頭からk=3のところで分けて、分けた二つをそれぞれreverse
	//[7,6,5,4,3,2,1] -> [7,6,5] [4,3,2,1]
	//[5,6,7] [1,2,3,4]

	//適切にswapするためにk mod len(nums),そうしないと[1,2] k=3の時にswapできない
	k = k % len(nums)

	//全体swap
	swap(0, len(nums)-1, nums)

	//kまでの要素をswap
	swap(0, k-1, nums[:k])

	//kからの要素をswap
	swap(k, len(nums)-1, nums[k:])

}

// @lc code=end


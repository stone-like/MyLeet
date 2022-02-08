/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start

//二分探索の問題、昇順ソートされた列が回転していて、ある性質を使う
//[1,2,3,4,5]を例として、

// [5,1,2,3,4]

// [4,5,1,2,3]

// [3,4,5,1,2]

// [2,3,4,5,1]
//この時にソートされている部分が二つに分けれているのがわかると思う。
//[3,4,5,1,2]であれば3~5,1~2,これを左ソート列、右ソート列とする

//どこでもでもいいけどとりあえず中央をMとして、端をL,Rとする
//このときにＭが左ソート列か右ソート列に含まれているかを知りたい
//知りたい理由は右ソート列は常に左ソート列より大きくなるため、左ソート列だけ見たいから
// M >= LであればMは左ソート列に入っているので、M+1をLとして見ていけばいいこれを使う
//こんな風にしてL,Rの範囲を狭めていき、もしL < Rとなれば終了
// L < Rとなった時にはソート済みということなのでLを取得すればいい

//Mの値は二度と見ないので、
//一番初めにMにminが来た時が危ないけど、必ずMの値をansに含めるようにすればいい
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func findMin(nums []int) int {

	left := 0
	right := len(nums) - 1

	ans := nums[left]

	for {
		if nums[left] <= nums[right] {
			ans = min(ans, nums[left])
			break
		}

		middle := (left + right) / 2
		ans = min(ans, nums[middle])
		if nums[middle] >= nums[left] {
			left = middle + 1
		} else {
			right = middle - 1
		}

	}

	return ans
}

// @lc code=end


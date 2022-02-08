/*
 * @lc app=leetcode id=1838 lang=golang
 *
 * [1838] Frequency of the Most Frequent Element
 */

// @lc code=start
//ソート&slidingWindow
//結構意図を読み取ることが難しいんだけど、やりたいことは
//[1,1,1,2,2,4] k=2で、
//l,r = 0の状態から、
// nums[r]*windowLen < total + kの条件でrを増やしていく、条件を満たさなかったらlを増やす
//何をやっているのかというと、windowを色々変えて、window内のすべてが同じ値になるようにしている
//  l   r
// [1,1,1,2,2,4]
//1*3 <= 3 + 2 なのでexpand(このwindowは全部同じ値にできる(もともと同じ値だけど))

//  l     r
// [1,1,1,2,2,4]
//2*4 <= 5+2 なのでleftMove(kを使ってもこのwindowは全部同じに出来ない)

//    l   r
// [1,1,1,2,2,4]
//2*3 <= 4+2 expand

//    l     r
// [1,1,1,2,2,4]
//2*4 <= 6+2 expand <--ここが最大(windowSizeが変わるごとにそのmaxをとればいい)

//    l       r
// [1,1,1,2,2,4]
//4*5 <= 10+2 leftMove..と続く

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxFrequency(nums []int, k int) int {

	sort.Ints(nums)

	maxFreq := -1
	right, left := 0, 0
	curTotal := 0

	for right < len(nums) {

		windowLen := right - left + 1

		if nums[right]*windowLen <= (curTotal+nums[right])+k {
			curTotal += nums[right]
			//rightMove
			right++
			//rightMoveの時のみwindow内の要素をすべて同じに出来るのでmaxしてよい
			maxFreq = max(maxFreq, windowLen)
			continue
		}

		//leftMove
		curTotal -= nums[left]
		left++

	}

	return maxFreq

}

// @lc code=end


/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
//ブルートフォースだとO(n^2)になるので、LとRをそれぞれ左端と右端において、ボトルネックとなるminになってる側->や<-に動かす、最初から左端と右端にポインタを置いているので見逃す心配はない、ボトルネックになるminを動かすので常に最大を探すようになっている

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	calcArea := func(height, length int) int {
		return height * length
	}

	for {
		if left == right {
			break
		}

		//leftの方が小さかったらleftを->
		if height[left] <= height[right] {
			maxArea = max(maxArea, calcArea(height[left], right-left))
			left++
		} else {
			maxArea = max(maxArea, calcArea(height[right], right-left))
			right--
		}

		//そうでないときはrightを<-

	}

	return maxArea
}

// @lc code=end


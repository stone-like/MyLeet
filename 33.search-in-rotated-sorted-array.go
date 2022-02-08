/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start

// [4,5,6,7,0,1,2]
//として、6=M、4=L,2=Rとする
//まずnums[M] >= nums[L]になっていればMはleftPortion,で違うならばrightPortionに入る

//この時targetの値によってleftPortionかrightPortionにを調べればいいかが変わる

//MがleftPortionに含まれるとき
//targetがMより大きい
//right側へ寄せる(l=mid+1)

//targetがM未満
//このときはleftPortionの一番左端(left)を調べて、target < leftならleftへright側へ
//target >= leftならleft側

//MがrightPortion(例えばM=1)
//targetがM未満
//left側(r = mid-1)へ

//targetがMより大きい
//この時はrightPortionの一番右端(right)を調べて right < targetならleft側へ
//target <= rightならright側へ

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for {
		if left > right {
			break
		}

		mid := (left + right) / 2

		if target == nums[mid] {
			return mid
		}

		if nums[mid] >= nums[left] {
			//midがLeftPortion
			if target > nums[mid] || target < nums[left] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			//midがRightPortion
			if target < nums[mid] || target > nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

// @lc code=end


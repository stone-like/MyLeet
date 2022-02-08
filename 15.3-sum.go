/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start

//ソートしたのちに左端の文字を一つ固定して2sumへ帰着
//target = 0で[-3,-3,0,3]みたいに-3が重複してしまう時のために
//i-1が同じ数字だったら使わないようにする、ソート済みなのでこのテクニックが使える
//ソートによる重複除去と一つ変数を固定することが本問のポイント

func twoSum(i, a, target int, nums []int) [][]int {
	var res [][]int

	left := i + 1
	right := len(nums) - 1

	for {
		if left >= right {
			break
		}

		twoSum := nums[left] + nums[right]

		// if twoSum == target {
		// 	res = append(res, []int{a, nums[left], nums[right]})
		// }

		if target > twoSum {
			left++
		} else if target < twoSum {
			right--
		} else {
			//ヒットしたとき
			res = append(res, []int{a, nums[left], nums[right]})
			left++
			//重複除去,right側の重複は今はどうでもよくて、左側の重複だけ気にしていけばいい,左の重複から消していけば同じ解にはならないし、いずれ->と進めていく内に右側の重複も消える
			for {
				if left >= right || nums[left] != nums[left-1] {
					break
				}
				left++
			}
		}
	}

	return res
}

func threeSum(nums []int) [][]int {
	var res [][]int

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i, v := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			//重複不可なので,このとき[-3,-3,0,3]みたいのでも-3みたいな重複するやつがある回答は最初の-3で作っているのでもう-3は要らない
			continue
		}

		//iを除いた箇所で2SUM
		//a+b+c = 0で今aを固定しているので b+c = -aとなるb,cを探す
		target := -1 * v
		twoSumRet := twoSum(i, v, target, nums)

		if twoSumRet != nil {
			res = append(res, twoSumRet...)
		}
	}

	return res
}

// @lc code=end


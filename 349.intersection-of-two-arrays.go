/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	num1Set := make(map[int]struct{})
	num2Set := make(map[int]struct{})
	result := make([]int, 0)

	for _, v := range nums1 {
		_, ok := num1Set[v]

		if ok {
			continue
		}

		num1Set[v] = struct{}{}
	}
	for _, v := range nums2 {
		_, ok := num2Set[v]

		if ok {
			continue
		}

		num2Set[v] = struct{}{}
	}

	for k1, _ := range num1Set {
		for k2, _ := range num2Set {
			if k1 == k2 {
				result = append(result, k1)
			}
		}
	}

	return result
}

// @lc code=end


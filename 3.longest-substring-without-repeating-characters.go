/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
//slideing windowsを使う、endを一個ずつ->にずらしていって、重複がでたら重複がなくなるまでstartをずらしていく
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	maxNum := 0
	left := 0
	count := 0

	wordMap := make(map[string]struct{})

	for right := 0; right < len(s); right++ {
		for {
			_, exists := wordMap[string(s[right])]

			if !exists {
				break
			}

			delete(wordMap, string(s[left]))

			left++
			count--
		}
		count++
		wordMap[string(s[right])] = struct{}{}
		maxNum = max(maxNum, count)
	}

	return maxNum

}

//下記は愚直解、これでもOKだったがもっと効率がいい方法を上に書く
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}

// 	return b
// }

// func lengthOfLongestSubstring(s string) int {
// 	if s == "" {
// 		return 0
// 	}

// 	maxNum := 0
// 	for i, r := range s {
// 		curStr := string(s[i])
// 		count := 1

// 		dict := map[string]struct{}{
// 			string(s[i]): {},
// 		}
// 		for j := i + 1; j < len(s); j++ {
// 			_, exist := dict[string(s[j])]

// 			if exist {
// 				break
// 			}

// 			curStr += string(r)
// 			count++
// 			dict[string(s[j])] = struct{}{}

// 		}

// 		maxNum = max(maxNum, count)
// 	}

// 	return maxNum

// }

// @lc code=end


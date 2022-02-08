/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */

// @lc code=start
//KMPAlgorthmによる解
//下記が詳しい
//https://www.youtube.com/watch?v=JoF0Z7nVSrA
//キーポイントとしてLongestPrefixSuffixを使う
//LPSとはprefixとsuffixが同じで一番長いLenのこと
//ex. AAだったらAとAで1,AAAだったらAAとAAでlen(AA) = 2
//AAAの時とかにAAAみたいに自分自身は含まない

func strStr(haystack string, needle string) int {

	if len(needle) == 0 {
		return 0
	}

	lps := make([]int, len(needle))

	prevLPS, ind := 0, 1

	for ind < len(needle) {
		if needle[prevLPS] == needle[ind] {
			//ここでやっていることは例えばAAAだったとして、最初のAの分のLPSは0で決まっているから二つ目のi=1から始める。
			//i=1のときは単純だからいいとして、i=2の時はなぜlps[ind]=2と即座にわかるかというと
			//i=1の時に needle[prevLPS](A) == needle[i](A)で既にAAとなっていることがわかり、前の二つは同じとi=2のときには分かっていて、i=1とi=2が同じならもうAAAしかないので即座にi=2のときのlpsは2とわかる
			lps[ind] = prevLPS + 1
			prevLPS += 1
			ind++
		} else if prevLPS == 0 {
			// A Bみたいなとき、この時は素直に0とするだけ
			lps[ind] = 0
			ind++
		} else {
			// CAACみたいに一致しなくなったときは現在のprevLPSから前にさかのぼって一致するところを確かめていく
			prevLPS = lps[prevLPS-1]

		}
	}

	i := 0
	j := 0
	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			if j == 0 {
				i++
			} else {
				j = lps[j-1]
			}
		}

		if j == len(needle) {
			return i - len(needle)
		}
	}

	return -1

}

// func strStr(haystack string, needle string) int {

// 	for i := 0; ; i++ {

// 		for j := 0; ; j++ {
// 			if j == len(needle) {
// 				return i
// 			}

// 			if i+j >= len(haystack) {
// 				return -1
// 			}

// 			if haystack[i+j] != needle[j] {
// 				break
// 			}
// 		}

// 	}

// 	return -1
// }

// @lc code=end


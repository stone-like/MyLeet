/*
 * @lc app=leetcode id=1888 lang=golang
 *
 * [1888] Minimum Number of Flips to Make the Binary String Alternating
 */

// @lc code=start

func generate(s string) (target1 string, target2 string) {
	//下記のやり方だとGoでは遅くなってしまうためstringBuilderを使う(strings.Repeatは内部でstringBuilderを使っている)
	// for i, _ := range s {

	// 	if i%2 == 0 {
	// 		target1 += "0"
	// 		target2 += "1"
	// 	} else {
	// 		target1 += "1"
	// 		target2 += "0"
	// 	}

	// }
	//strings
	target1 = "01"
	target2 = "10"

	repeatNum := len(s) / 2

	target1 = strings.Repeat(target1, repeatNum)
	target2 = strings.Repeat(target2, repeatNum)

	return
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

//メインアイデアとしてrotateするのをそのままsを足して表現してしまう
//それでsliding windowsとして解く

func minFlips(s string) int {

	startRotateIndex := len(s)

	s = s + s

	target1, target2 := generate(s)

	res := len(s)
	diff1, diff2 := 0, 0
	left := 0

	for right := 0; right < len(s); right++ {
		//rightはrotateに入る前も入った後も変わらず
		if s[right] != target1[right] {
			diff1++
		}

		if s[right] != target2[right] {
			diff2++
		}

		//s + sで最初のsが終わって二つ目のsに入ったときはrotateの表現に入った時
		//leftは一回目のsのindexに対応させる
		if (right - left + 1) > startRotateIndex {
			if s[left] != target1[left] {
				diff1--
			}

			if s[left] != target2[left] {
				diff2--
			}

			left++
		}

		//rotateに入る直前+入って以降毎回チェック
		if (right - left + 1) == startRotateIndex {
			res = min(res, min(diff1, diff2))
		}

	}

	return res
}

// @lc code=end


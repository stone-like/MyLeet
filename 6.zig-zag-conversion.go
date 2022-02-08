/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {

	if len(s) == 1 {
		return s
	}

	if numRows == 1 {
		return s
	}

	ans := ""

	jump := numRows + (numRows - 2)

	curIndex := 0
	curRows := 0

	for {
		if len(ans) == len(s) {
			break
		}

		if curIndex > len(s)-1 {
			backIndex := curIndex - 2*curRows
			if curRows != 0 && curRows != numRows-1 && backIndex <= len(s)-1 {
				//backした文字を入れる必要がある
				backIndex := curIndex - 2*curRows
				ans += string(s[backIndex])
			}
			curRows++
			curIndex = curRows
			continue
		}

		if curRows != 0 && curRows != numRows-1 {
			//backした文字を入れる必要がある
			backIndex := curIndex - 2*curRows
			if 0 <= backIndex {
				ans += string(s[backIndex])
			}

		}

		ans += string(s[curIndex])

		curIndex = curIndex + jump

		//P    i
		//A  L S
		//Y A  H
		//P    i
		//上のように会ってP -> iとなめていく
		//P -> iにジャンプするために必要なjumpIndexは
		//numRows + (numsRows - start+end)
		//PAYPの最初と最後のPとPはjumpIndexが楽で
		//中盤のjumpIndexも最初と最後のrowと同じようなジャンプでいいけど＋して奇数行だったらその行数*2戻ったやつを先に＋する

	}

	return ans
}

// func dup(str string, num int) string {
// 	temp := ""
// 	for i := 0; i < num; i++ {
// 		temp += str
// 	}

// 	return temp
// }

// func convert(s string, numRows int) string {
// 	//3 -> 4
// 	//4 -> 6
// 	//5 -> 8
// 	//len + (len - start+end)が1ブロック

// 	if len(s) == 1 {
// 		return s
// 	}

// 	if numRows == 1 {
// 		return s
// 	}

// 	//一ブロック当たりの処理できる文字の個数
// 	block := numRows + (numRows - 2)

// 	//s=14,numRows=3だと一ブロックで処理できる文字は4
// 	// 14 /3+1 = 3.4..で切り上げして4ブロックいる
// 	//一ブロックはindex0,1で構成される
// 	//P
// 	//A P
// 	//Y
// 	//0 1
// 	//なので1ブロック当たりのindex数がssBlockを出して
// 	//ブロック数*ssBlockがssの個数

// 	//何ブロックいるか
// 	n := int(math.Ceil(float64(len(s)) / float64(block)))

// 	//1ブロック当たりのindex数
// 	ssBlock := 1 + numRows - 2

// 	ss := make([]string, n*ssBlock)

// 	for i := 0; i < n*ssBlock; i++ {
// 		ss[i] = dup("+", numRows)
// 		// if i%ssBlock != 0 {
// 		// 	ss[i] = dup("+", numRows)
// 		// }
// 	}

// 	ssInd := 0
// 	targetZero := 0

// 	for _, r := range s {
// 		//numRowが6だったら0,1,2が1ブロック、3,4,5がまた一ブロックとssのindexは3つ単位で1ブロック,1blockはがどのくらいのindexの単位かはssBlock
// 		//index0の時は普通に詰めて、
// 		//1の時はnumRows-1-1のとこだけ
// 		//2のときはnumRows-1-2のとこだけ

// 		target := ssInd % ssBlock

// 		if target == 0 {

// 			temp := ss[ssInd]
// 			ss[ssInd] = temp[:targetZero] + string(r) + temp[targetZero+1:]
// 			targetZero++

// 			if targetZero == numRows {
// 				targetZero = 0
// 				ssInd++
// 			}
// 		} else {
// 			temp := ss[ssInd]
// 			targetIndex := numRows - 1 - target
// 			ss[ssInd] = temp[:targetIndex] + string(r) + temp[targetIndex+1:]
// 			ssInd++
// 		}

// 	}

// 	ans := ""

// 	for i := 0; i < numRows; i++ {
// 		for j := 0; j < len(ss); j++ {
// 			r := ss[j][i]
// 			if r == '+' {
// 				continue
// 			}

// 			ans += string(r)
// 		}
// 	}

// 	return ans
// }

// @lc code=end


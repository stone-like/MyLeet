/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start

func peek(p string, index int) byte {
	return p[index+1]
}

func isDot(p string, index int) bool {
	return p[index] == '.'
}

func isRemainingAllAsterisk(p string, curP int) bool {
	//bba*c*d*で、s=bb、curStr=bbのときにヒットさせたい

	//bba*cはだめなパターン

	curIndex := curP
	isAllAsterRisk := true
	for {
		if !isPValidLen(p, curIndex+1) {
			if p[curIndex] != '*' {
				isAllAsterRisk = false
			}
			break
		}

		if peek(p, curIndex) == '*' {
			if len(p)-1 == curIndex+1 {
				break
			}

			curIndex += 2
		} else {
			isAllAsterRisk = false
			break
		}

	}

	return isAllAsterRisk
}

func isValidPState(p string, curP int) bool {
	//もしaaaaa , a*みたいに*のときだったらcurPの次が*でそれでpが終わりならOk
	//*でないならpが終わりまで来ていればOK

	if isPValidLen(p, curP+1) && isRemainingAllAsterisk(p, curP) {
		return true
	}

	return len(p) == curP
}

func isPValidLen(p string, curP int) bool {
	return curP <= len(p)-1
}

func dfs(s, p, curStr string, curS, curP int) bool {

	//一応ここまできたらcurSまでsとcurStrは一致しているはず
	//sが最後まで来てかつ、pも最後まで来ていればOK
	// if len(s) == curS && len(p) == curP {
	// 	return true
	// }

	//len(p)条件はaaaaa , a*bのときにcurStrがaaaaaとなっていても実際pはまだbが余っているからその処理
	//curPの次が*でそれでpが終わりなのであればOK
	if isValidPState(p, curP) && s == curStr {
		return true
	}

	if len(s) <= curS {
		//ccccca と .*aで.*をcccccaの数以上にやった時
		//ccaもccccaも.*aでヒットさせなくてはいけない
		return false
	}

	//peekして*だったら*を使う時使わないときで分ける
	if curP+1 <= len(p)-1 && peek(p, curP) == '*' {
		//この場合curPIndexは進めない(*なので)、進めないままでsをオーバーしても上でfalseになるのでOK
		ret1 := false
		ret2 := false
		//targetを使う場合->equalもしくはDotか確認
		if p[curP] == s[curS] || isDot(p, curP) {
			ret1 = dfs(s, p, curStr+string(s[curS]), curS+1, curP)
		}

		//使わず*を0個とする場合
		if curP+2 <= len(p)-1 {
			//curP+3しているのはca,c*aとして今index1としてPの方もaが終わったことにするため
			ret2 = dfs(s, p, curStr, curS, curP+2)
		}

		if ret1 || ret2 {
			return true
		}
	} else {

		if !isPValidLen(p, curP) {
			return false
		}

		if p[curP] == s[curS] || isDot(p, curP) {
			return dfs(s, p, curStr+string(s[curS]), curS+1, curP+1)
		}

		return false

	}

	return false

}

func isMatch(s string, p string) bool {
	//バックトラックかな？例えばc*b*aとあったらcを0~N回使い,sとマッチしなくなったらその枝は終わり

	return dfs(s, p, "", 0, 0)
}

// @lc code=end


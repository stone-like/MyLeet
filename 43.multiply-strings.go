/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */

// @lc code=start
func reverse(s string) string {

	temp := ""
	for i := len(s) - 1; i >= 0; i-- {
		temp += string(s[i])
	}
	return temp
}

func multiply(num1 string, num2 string) string {

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	addSlice := make([]int, len(num1)+len(num2))

	rev1 := reverse(num1)
	rev2 := reverse(num2)

	//  num1
	//× num2
	//の計算
	for i, v1 := range rev1 {
		for j, v2 := range rev2 {
			ret := int(v1-'0') * int(v2-'0')
			addSlice[i+j] += ret                  //最初に全部足してあげる
			addSlice[i+j+1] += addSlice[i+j] / 10 //carryを出す
			addSlice[i+j] = addSlice[i+j] % 10    //carryした分を消して一桁にする
		}
	}

	temp := ""
	isNumber := false
	for i := len(addSlice) - 1; i >= 0; i-- {
		if addSlice[i] == 0 && !isNumber {
			continue
		}

		temp += strconv.Itoa(addSlice[i])
		isNumber = true

	}

	return temp

}

// @lc code=end


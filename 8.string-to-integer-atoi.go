/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start

var int32Inf = 2147483648 - 1
var int32MinusInf = -2147483648

func Pow(a, b float64) float64 {
	return math.Pow(a, b)
}

func OwnAtoi(s string) float64 {
	var temp float64
	minus := false

	if s[0] == '-' {
		minus = true
		s = s[1:]
	}

	length := len(s) - 1

	curPos := length

	for i := length; i >= 0; i-- {

		if s[i] == '.' {
			temp = temp / Pow(10, float64(length-i))
			curPos = i - 1
			continue
		}
		num := s[i] - '0'
		multi := Pow(10, float64(curPos-i))
		temp += float64(multi) * float64(num)
	}

	if minus {
		return temp * float64(-1)
	}

	return temp
}
func Parser(s string) int {
	minus := false
	plus := false
	tempStr := ""
	isAlreadyHasWord := false
	for _, r := range s {
		if r == '-' {
			if len(tempStr) != 0 {
				break
			}

			if plus || minus {
				return 0
			}

			minus = true
			continue
		}

		if r == '+' {
			if len(tempStr) != 0 {
				break
			}

			if plus || minus {
				return 0
			}
			plus = true
			continue
		}
		if unicode.IsSpace(r) {
			//"+413  "はOKだけど
			//"+  413"はだめ
			if len(tempStr) != 0 {
				break
			}

			if len(tempStr) == 0 && (plus || minus) {
				return 0
			}
			continue
		}

		if '0' <= r && r <= '9' || r == '.' {
			if isAlreadyHasWord {
				break
			}
			tempStr += string(r)
			continue
		}

		isAlreadyHasWord = true

	}

	if plus && minus {
		return 0
	}

	if tempStr == "" {
		return 0
	}

	if minus {
		tempStr = "-" + tempStr
	}

	intNUM := OwnAtoi(tempStr)

	if intNUM < float64(int32MinusInf) {
		intNUM = float64(int32MinusInf)
	}

	if float64(int32Inf) < intNUM {
		intNUM = float64(int32Inf)
	}

	return int(intNUM)
}

func myAtoi(s string) int {
	return Parser(s)
}

// @lc code=end


/*
 * @lc app=leetcode id=65 lang=golang
 *
 * [65] Valid Number
 */

// @lc code=start
const (
	numFlag  uint = 1 << iota //0001
	dotFlag                   //0010
	eFlag                     //0100
	signFlag                  //1000
)

func flagIsActive(state, target uint) bool {
	//例えばstate = 0000でtarget = 0110(dot or E)
	//もしDotかEFlagが立っていればその部分は1&1で1になるので0にはならない
	return state&target != 0
}

func setFlag(state, target uint) uint {
	return state | target
}

func removeFlag(state, target uint) uint {
	return state &^ target
}

func isNumber(s string) bool {
	str := s
	state := uint(0)
	// numFlag, dotFlag, eFlag, plusOrMinusFlag := false, false, false, false

	processDot := func(i int) bool {
		if flagIsActive(state, dotFlag|eFlag) {
			return false
		}

		//終了時にnumFlagがfalseかつ、dotFlag=trueはダメ->".1","1."はOK,"."はダメ

		if i == len(str)-1 && !flagIsActive(state, numFlag) {
			return false
		}

		return true
	}

	processE := func(i int) bool {
		if flagIsActive(state, eFlag) {
			return false
		}

		//eの前にnumFlagがないとダメ
		if !flagIsActive(state, numFlag) {
			return false
		}

		//最後がe,Eだとだめ
		if i == len(str)-1 {
			return false
		}

		return true
	}

	processSign := func(i int) bool {
		if flagIsActive(state, signFlag|numFlag|dotFlag) {
			return false
		}

		//終了時にnumFlagがfalseかつplusOrMinusFlagがtrueはだめ
		if i == len(str)-1 && !flagIsActive(state, numFlag) {
			return false
		}

		return true
	}

	for i := 0; i < len(str); i++ {
		target := str[i]
		switch {
		case '0' <= target && target <= '9':
			state = setFlag(state, numFlag)
		case target == '.':
			{
				if !processDot(i) {
					return false
				}

				state = setFlag(state, dotFlag)

			}
		case target == 'e' || target == 'E':
			{
				if !processE(i) {
					return false
				}

				state = removeFlag(state, numFlag|dotFlag|signFlag)
				state = setFlag(state, eFlag)

			}
		case target == '+' || target == '-':
			{
				if !processSign(i) {
					return false
				}
				state = setFlag(state, signFlag)

			}
		default:
			return false
		}
	}

	return true
}

// @lc code=end


/*
 * @lc app=leetcode id=779 lang=golang
 *
 * [779] K-th Symbol in Grammar
 */

// @lc code=start

//愚直に作成するのではなく下記の対応関係を使う
// 0(k=1)    1(k=2)            row N-1
//0(k=1) 1(k=2) 1(k=3) 0(k=4)  row N

//子から親を求めるにはkthGrammer(N-1,[k/2]) []は切り上げ

func kthGrammar(n int, k int) int {

	if n == 1 {
		return 0
	}

	//一階層ずつ上がっていく
	parent := kthGrammar(n-1, int(math.Ceil(float64(k)/float64(2))))

	isOdd := k%2 != 0

	if parent == 0 {
		if isOdd {
			return 0
		}

		return 1
	}

	//parentが1の時
	if isOdd {
		return 1
	}

	return 0

}

// var convert = map[string]string{
// 	"0": "01",
// 	"1": "10",
// }

// func Traverse(n int, curStr string) string {

// 	convertedStr := convert[curStr]
// 	if n == 0 {
// 		return convertedStr
// 	}

// 	temp := ""
// 	for _, r := range convertedStr {
// 		temp += Traverse(n-1, string(r))
// 	}

// 	return temp
// }

// func kthGrammar(n int, k int) int {

// 	s := "0"

// 	if n == 1 {
// 		return 0
// 	}

// 	return int(Traverse(n-2, s)[k-1] - '0')
// }

// @lc code=end


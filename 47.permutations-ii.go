/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
//BackTrackを使ったパターン
//[1,1,2]みたいに元が重複している数字を持つ配列だと結果も重複してしまうので、あらかじめmapでcount形式にしてやってあげればいい

func copy(s []int) []int {
	temp := make([]int, len(s))

	for i, v := range s {
		temp[i] = v
	}

	return temp
}

func dfs(res *[][]int, m map[int]int, current, nums []int) {

	if len(current) == len(nums) {
		*res = append(*res, copy(current))
		return
	}

	for k, count := range m {

		if count == 0 {
			continue
		}

		m[k] = count - 1
		current = append(current, k)
		dfs(res, m, current, nums)
		current = current[:len(current)-1]
		m[k] = count
	}
}

func permuteUnique(nums []int) [][]int {

	m := make(map[int]int)

	var res [][]int

	for _, v := range nums {
		_, exists := m[v]

		if !exists {
			m[v] = 1
		} else {
			m[v]++
		}
	}

	dfs(&res, m, []int{}, nums)

	return res
}

//下記は結構ごり押し
// func Permute(nums []int) [][]int {
// 	n := factorial(len(nums)) //配列の数に応じた総パターン数,len=3なら3!=6
// 	ret := make([][]int, 0, n)
// 	heapPermute(len(nums), makeCopy(nums), &ret)
// 	return ret
// }

// //quickPermの方がちょっと良いくらいだが、heapPermuteの方が簡単なのでこっちを使う

// func alreadyHas(ret *[][]int, pat []int) bool {

// 	for _, s := range *ret {
// 		has := true

// 		for i := 0; i < len(s); i++ {
// 			if pat[i] != s[i] {
// 				has = false
// 			}
// 		}

// 		if has {
// 			return true
// 		}
// 	}

// 	return false
// }

// //heapAlgorithmについては下記が詳しい
// //https://medium.com/sodalabs/heaps-algorithm-fun-observation-4986a188a80
// func heapPermute(n int, pat []int, ret *[][]int) {
// 	if n == 1 {

// 		if alreadyHas(ret, pat) {
// 			return
// 		}
// 		*ret = append(*ret, makeCopy(pat))
// 	} else {
// 		//[1,2,3]の例だったら
// 		//[1,2,3]レベルのpermを3回する
// 		//n==1に行ったときに[1,2,3]をretにpushしたあとn==2に返ってきて
// 		//[1,2,3]の3を固定した[1,2]レベルのpermを2回する
// 		//これで3*2 = 6

// 		//全体の流れを書くと
// 		//[1,2,3]
// 		//[2,1,3]
// 		//[3,1,2]
// 		//[1,3,2]
// 		//[2,3,1]
// 		//[3,2,1]

// 		for i := 0; i < n; i++ {
// 			heapPermute(n-1, pat, ret)

// 			//ここの条件は
// 			//n=現在見ている部分での配列のlenで
// 			//n=偶数なら nの範囲だけで考えてmostleftをnの最後に
// 			//[1,2,3]でn=2で[1,2]レベルの時ならleftmostの1をnの最後に

// 			//n=奇数なら配列の最初と最後をswap
// 			if n%2 == 0 {
// 				pat[i], pat[n-1] = pat[n-1], pat[i]
// 			} else {
// 				pat[0], pat[n-1] = pat[n-1], pat[0]
// 			}
// 		}
// 	}
// }

// //QuickPerm
// func quickPerm(nums []int, ret *[][]int) {
// 	*ret = append(*ret, makeCopy(nums))

// 	n := len(nums)
// 	p := make([]int, n+1)
// 	for i := 0; i < n+1; i++ {
// 		p[i] = i
// 	}
// 	for i := 1; i < n; {
// 		p[i]--
// 		j := 0
// 		if i%2 == 1 {
// 			j = p[i]
// 		}

// 		nums[i], nums[j] = nums[j], nums[i]
// 		*ret = append(*ret, makeCopy(nums))
// 		for i = 1; p[i] == 0; i++ {
// 			p[i] = i
// 		}
// 	}
// }

// func factorial(n int) int {
// 	ret := 1
// 	for i := 2; i <= n; i++ {
// 		ret *= i
// 	}
// 	return ret
// }

// func makeCopy(nums []int) []int {
// 	return append([]int{}, nums...)
// }

// func permuteUnique(nums []int) [][]int {
// 	return Permute(nums)
// }

// @lc code=end


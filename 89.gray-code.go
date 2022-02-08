/*
 * @lc app=leetcode id=89 lang=golang
 *
 * [89] Gray Code
 */

// @lc code=start
//アイデアとしては、N=0=0,N=1(2^0)..とやっていき、それぞれ前のNで生成されたものにNを足していく
//例えばN=0なら0のみ生成で、N=1の時に0+1=1で生成、前のNで生成された奴も引き継いで0,1が生成
//N=2(2^1)なら0+2 = 2,1+2=3で0,1,2,3
//N=3(2^2)なら4,5,6,7が生成

//graycodeの本旨である1箇所だけ変わっていくという条件は
//N=2なら0,1に2を足すところで、1,0と足していけばいい0,1,3,2となり1箇所のみの変化

func grayCode(n int) []int {

	var ans []int
	ans = append(ans, 0)

	if n == 0 {
		return ans
	}
	ans = append(ans, 1)
	// ans[0] = 0
	// ans[1] = 1

	curNum := 1
	for i := 2; i <= n; i++ {
		curNum *= 2
		ansLen := len(ans)
		for j := ansLen - 1; j >= 0; j-- {
			ans = append(ans, ans[j]+curNum)
		}
	}

	return ans

}

// @lc code=end


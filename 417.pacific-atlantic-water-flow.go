/*
 * @lc app=leetcode id=417 lang=golang
 *
 * [417] Pacific Atlantic Water Flow
 */

// @lc code=start
//アイデアとしてはfor m for nのループですべてに対してdfsをして、
//        Ocean
// Ocean　1 2 3 とあって、1の結果を2,3と伝播させていく方式ではTLEしてしまう
//なので逆にy=0の行、x=0の列はPacificOceanに到達しているとして、そこからどの(y,x)がPacificに到達できるか終わりから見ていく
//y=mの行、x=nの列はAtlanticOceanに到達しているとして、こっちも同じようにする
//最後に for m for nのループでpacificに到達かつatlanticに到達しているものを探せばいい
var dirs = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func isOnBoard(x, y int, heights [][]int) bool {
	return x >= 0 && x < len(heights[0]) && y >= 0 && y < len(heights)
}

func dfs(x, y int, visited *[][]bool, heights [][]int) {

	(*visited)[y][x] = true
	//ここのvisited= trueということはpacificから始めればpcificまで到達できるということ,atlanticもそう

	for _, dir := range dirs {
		dx := x + dir[0]
		dy := y + dir[1]

		if !isOnBoard(dx, dy, heights) {
			continue
		}

		//Ocean側からやっているので、heightが高い方には行かないのではなく、heightが高い方にしか進めない
		if heights[y][x] > heights[dy][dx] {
			continue
		}

		if (*visited)[dy][dx] {
			continue
		}

		dfs(dx, dy, visited, heights)

	}

}

func pacificAtlantic(heights [][]int) [][]int {

	if len(heights) == 0 || len(heights[0]) == 0 {
		return nil
	}

	m, n := len(heights), len(heights[0])
	ans := make([][]int, 0)

	pacific, atlantic := make([][]bool, m), make([][]bool, m)

	for i := 0; i < m; i++ {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		//右端から
		dfs(0, i, &pacific, heights)
		//左端から
		dfs(n-1, i, &atlantic, heights)
	}

	for j := 0; j < n; j++ {
		//上端から
		dfs(j, 0, &pacific, heights)
		//下端から
		dfs(j, m-1, &atlantic, heights)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if pacific[i][j] && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}

	return ans
}

// @lc code=end


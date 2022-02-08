/*
 * @lc app=leetcode id=1466 lang=golang
 *
 * [1466] Reorder Routes to Make All Paths Lead to the City Zero
 */

// @lc code=start
//問題の前提条件として、n個の街にn-1の道で、onewayなのでループがなく、connectedであることがわかる、n個でn-1個の道だとループは作れない
//有効グラフをedgeとneighborを使って作り、
//もしneighborであってedgeがあるのなら、0には到達しないのでcount+1
//ex. 0 -> 1 -> 3
//edgeがあっても無くてもneighborなら辿るので
//0 -> 1 -> 3 <- 2となっていても3から2をたどる
//よって一応すべてのNodeは0からたどれることになる、edgeが無いときは辿らないとしていたため回答が厳しかった

type edge struct {
	to   int
	from int
}

func isEdgeExists(from, to int, edgeMap map[edge]struct{}) (ok bool) {
	edge := edge{from: from, to: to}
	_, ok = edgeMap[edge]

	return
}

func dfs(visited *[]bool, edgeMap map[edge]struct{}, neighbor map[int][]int, curNode int, count *int) {

	for _, nei := range neighbor[curNode] {
		if (*visited)[nei] {
			continue
		}

		if isEdgeExists(curNode, nei, edgeMap) {
			(*count) += 1
		}
		(*visited)[nei] = true
		dfs(visited, edgeMap, neighbor, nei, count)

	}

}

func minReorder(n int, connections [][]int) int {
	edgeMap := make(map[edge]struct{})

	for _, pair := range connections {
		edgeMap[edge{from: pair[0], to: pair[1]}] = struct{}{}
	}

	neighbor := make(map[int][]int)

	for _, pair := range connections {
		nei1 := pair[0]
		nei2 := pair[1]

		neighbor[nei1] = append(neighbor[nei1], nei2)
		neighbor[nei2] = append(neighbor[nei2], nei1)
	}

	count := 0

	visited := make([]bool, n)
	visited[0] = true
	dfs(&visited, edgeMap, neighbor, 0, &count)
	return count
}

// @lc code=end


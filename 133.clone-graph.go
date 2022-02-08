/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */

// @lc code=start

//dfsで良いが今回の肝は無向グラフなので1 <->2とあってdfsをしてしまう時に無限ループに陥ってしまうのをどう防ぐか
//解決策としてmapでOldNodeとCloneNodeの対応を使う
//1 -> 2に見るときに1 -> Clone 1の対応を作ってから2に行く
//2ではneightborとして1を見るけど、mapに1の値があればdfsせずに1をmapから取り出せばよい
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func dfs(node *Node, m map[int]*Node) *Node {

	clone := &Node{
		Val: node.Val,
	}

	_, exists := m[node.Val]
	if !exists {
		m[node.Val] = clone
	}
	for _, v := range node.Neighbors {

		neighborClone, exists := m[v.Val]
		if !exists {
			neighborClone = dfs(v, m)
		}
		clone.Neighbors = append(clone.Neighbors, neighborClone)
	}
	return clone
}

func cloneGraph(node *Node) *Node {

	if node == nil {
		return nil
	}
	m := make(map[int]*Node)

	return dfs(node, m)
}

// @lc code=end


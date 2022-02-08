/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func dfs(head *Node, m map[*Node]int, m2IndKey map[int]*Node, index int) *Node {

	if head == nil {
		return nil
	}

	//一時的にrandomは元のheadのやつを入れておく(後で変える)
	node := &Node{
		Val:    head.Val,
		Random: head.Random,
	}

	next := dfs(head.Next, m, m2IndKey, index+1)
	node.Next = next

	m[head] = index        //こっちは古いNodeで比較
	m2IndKey[index] = node //こっちは新しいnode

	return node
}
func copyRandomList(head *Node) *Node {
	//新しく作ったノードを保存してランダムに備える
	m := make(map[*Node]int)
	m2IndKey := make(map[int]*Node)
	newNode := dfs(head, m, m2IndKey, 0)

	head = newNode
	cur := head
	for cur != nil {
		if cur.Random != nil {
			ind := m[cur.Random] //古いNodeで比較
			node := m2IndKey[ind]
			cur.Random = node
		}
		cur = cur.Next
	}

	return newNode
}

// @lc code=end


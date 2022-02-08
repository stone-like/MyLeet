/*
 * @lc app=leetcode id=669 lang=golang
 *
 * [669] Trim a Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//もしくは*TreeNodeを返す方向で作成
//やっぱり値を返した方が参照透過的だしわかりやすいし好き
func trimBST(root *TreeNode, low int, high int) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Val < low {
		root = trimBST(root.Right, low, high)
		return root
	}

	if high < root.Val {
		root = trimBST(root.Left, low, high)
		return root
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	return root

}

//ポインタのポインタを使っているのは引数にポインタをそのまま渡すと、ポインタがコピーされてしまうので
//もちろんコピーされたポインタも同じ場所を参照しているから普通の用途では問題ないけど
//今回はポインタが指す値ではなく、ポインタの指す場所を変更したかったため、ポインタがコピーされてはいけなかった
// func search(node **TreeNode, low, high int) {

// 	no := *node

// 	if no == nil {
// 		return
// 	}

// 	if high < no.Val {
// 		*node = no.Left
// 		search(node, low, high)
// 		return
// 	}

// 	if no.Val < low {
// 		*node = no.Right
// 		search(node, low, high)
// 		return
// 	}
// 	search(&no.Right, low, high)
// 	search(&no.Left, low, high)

// }
// func trimBST(root *TreeNode, low int, high int) *TreeNode {

// 	search(&root, low, high)

// 	return root

// }

// @lc code=end


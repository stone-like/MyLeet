/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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

//[1,2,3,4,5,null,null]
//を例として
//  Inorder (Left, Root, Right) : 4 2 5 | 1 | 3
//  Preorder (Root, Left, Right) : 1 | 2 4 5 | 3
//  Postorder (Left, Right, Root) : 4 5 2 | 3 | 1
//重要なのは回り方の違いで、root,left,rightの個数自体は変わらないので
//例えばpreorderの最初は必ずrootなので、この値と同じindexをInorderから見つけるとRootより<-はLeftなのでrootのindexが分かればleftの個数とrightの個数がわかる
//なのでここからpreorderのleft部分とright部分がわかる
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	node := &TreeNode{
		Val: preorder[0],
	}

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == node.Val {

			node.Left = buildTree(preorder[1:i+1], inorder[:i])
			node.Right = buildTree(preorder[i+1:], inorder[i+1:])

		}
	}

	return node
}

// @lc code=end


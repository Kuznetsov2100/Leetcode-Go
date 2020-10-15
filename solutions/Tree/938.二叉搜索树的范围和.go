/*
 * @lc app=leetcode.cn id=938 lang=golang
 *
 * [938] 二叉搜索树的范围和
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
func rangeSumBST(root *TreeNode, L int, R int) int {
	var res int
	if root == nil {
		return res
	}
	var inorder func(x *TreeNode)
	inorder = func(x *TreeNode)	{
		if x == nil {
			return
		}
		inorder(x.Left)
		if x.Val >= L && x.Val <= R {
			res += x.Val
		}
		inorder(x.Right)
	}
	inorder(root)
	return res
	

}
// @lc code=end


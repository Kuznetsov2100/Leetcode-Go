/*
 * @lc app=leetcode.cn id=965 lang=golang
 *
 * [965] 单值二叉树
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
func isUnivalTree(root *TreeNode) bool {
	res := true
	var preorder func(x *TreeNode)
	preorder = func(x *TreeNode) {
		if x == nil {
			return
		}
		if x.Val != root.Val {
			res = false
			return
		}
		preorder(x.Left)
		preorder(x.Right)
	}
	preorder(root)
	return res
}
// @lc code=end


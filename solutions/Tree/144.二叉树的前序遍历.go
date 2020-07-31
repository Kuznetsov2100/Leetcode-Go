/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var preorder func(x *TreeNode)
	preorder = func(x *TreeNode) {
		if x == nil {
			return
		}
		res = append(res, x.Val)
		preorder(x.Left)
		preorder(x.Right)
	}
	preorder(root)
	return res
}
// @lc code=end


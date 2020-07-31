/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
	var res []int
	var postorder func(x *TreeNode)
	postorder = func(x *TreeNode) {
		if x == nil {
			return
		}
		postorder(x.Left)
		postorder(x.Right)
		res = append(res, x.Val)
	}
	postorder(root)
	return res
}
// @lc code=end


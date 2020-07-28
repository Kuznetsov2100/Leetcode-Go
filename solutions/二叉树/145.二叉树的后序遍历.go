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
	if root == nil {
		return nil
	}
	var res []int
	postorder(root, &res)
	return res
}

func postorder(x *TreeNode, res *[]int) *TreeNode {
	if x == nil {
		return nil
	}
	postorder(x.Left, res)
	postorder(x.Right, res)
	*res = append(*res, x.Val)
	return x
}
// @lc code=end


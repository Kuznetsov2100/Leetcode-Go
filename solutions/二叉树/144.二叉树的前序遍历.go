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
	if root == nil {
		return nil
	}
	var res []int
	preorder(root, &res)
	return res
}

func preorder(x *TreeNode, res *[]int) *TreeNode {
	if x == nil {
		return nil
	}
	*res = append(*res, x.Val)
	preorder(x.Left, res)
	preorder(x.Right, res)
	return x
}
// @lc code=end


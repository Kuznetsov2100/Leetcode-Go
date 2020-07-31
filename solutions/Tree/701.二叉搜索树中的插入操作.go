/*
 * @lc app=leetcode.cn id=701 lang=golang
 *
 * [701] 二叉搜索树中的插入操作
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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 涉及到修改操作,函数就要返回*TreeNode类型，并且对递归调用的返回值进行接收
	if root == nil {
		return &TreeNode{Val:val}
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
// @lc code=end


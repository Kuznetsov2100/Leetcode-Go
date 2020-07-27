/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int	
	inorder(root, &res)
	return res
}

// The pointer to a slice is indispensable when the function is going to modify the structure,
// the size, or the location in memory of the slice and every change should to be visible
// to those who call the function
func inorder(x *TreeNode, res *[]int) {
	if x == nil {
		return
	}
	inorder(x.Left, res)
	*res = append(*res, x.Val)
	inorder(x.Right, res)
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(max(diameterOfBinaryTree(root.Left),diameterOfBinaryTree(root.Right)),depth(root.Left) + depth(root.Right))
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return max(depth(root.Left),depth(root.Right))+1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end


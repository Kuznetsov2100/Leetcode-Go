/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := countLevel(root.Left)
	right := countLevel(root.Right)
	if left == right { // 左子树和右子树高度相同，表明左子树是满二叉树，左子树的节点数为2^left-1,加上根节点，然后递归计算右子树的节点数
		return countNodes(root.Right) + (1 << left)
	}
	// 高度不同，表明右子树是满二叉树，右子树的节点数为2^right-1，加上根节点，然后递归计算左子树的节点数
	return countNodes(root.Left) + (1 << right)
}

// 计算完全二叉树的高度
func countLevel(node *TreeNode) int {
	level := 0
	for node != nil {
		level++
		node = node.Left
	}
	return level
}
// @lc code=end


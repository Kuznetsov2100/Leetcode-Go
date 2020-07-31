/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
 import "math"
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(float64(maxDepth(root.Left)-maxDepth(root.Right))) > 1.0 {
		return false
	} 
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return max(maxDepth(root.Left)+1,maxDepth(root.Right)+1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end


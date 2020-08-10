/*
 * @lc app=leetcode.cn id=671 lang=golang
 *
 * [671] 二叉树中第二小的节点
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
func findSecondMinimumValue(root *TreeNode) int {
	// 根据题意，根节点一定是最小的元素，helper函数的作用是找到比根节点大的元素，即为第二小的元素
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		if node.Val > root.Val {
			return node.Val
		}
		left := helper(node.Left)
		right := helper(node.Right)
		if left == -1 {
			return right
		}
		if right == -1 {
			return left
		}
		return func(x, y int) int {
			if x < y {
				return x
			}
			return y
		}(left,right)
	}
	return helper(root)
}
// @lc code=end


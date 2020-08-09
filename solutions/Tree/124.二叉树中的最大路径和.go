/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
func maxPathSum(root *TreeNode) int {
	res := -1 << 63
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(0,dfs(node.Left))
		right := max(0, dfs(node.Right))
		res = max(res, node.Val + left + right)
		return node.Val + max(left,right)
	}
	dfs(root)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end


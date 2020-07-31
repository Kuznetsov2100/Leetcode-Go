/*
 * @lc app=leetcode.cn id=563 lang=golang
 *
 * [563] 二叉树的坡度
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
func findTilt(root *TreeNode) int {
	res := 0
	// 后序遍历求子树和，顺便求出总坡度
	var nodeSum func(x *TreeNode) int
	nodeSum = func(x *TreeNode) int {
		if x == nil {
			return 0
		}
		left := nodeSum(x.Left)
		right := nodeSum(x.Right)
		res += abs(left-right)
		return x.Val+left+right
	}
	nodeSum(root)
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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
func getMinimumDifference(root *TreeNode) int {
	minvalue := 1 << 63 -1 
	var pre *TreeNode
	var inorder func(x *TreeNode)
	inorder = func(x *TreeNode)	{
		if x == nil {
			return
		}
		inorder(x.Left)
		if pre != nil {
			minvalue = min(minvalue,x.Val-pre.Val)
		}
		pre = x
		inorder(x.Right)
	}
	inorder(root)
	return minvalue
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end


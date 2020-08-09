/*
 * @lc app=leetcode.cn id=687 lang=golang
 *
 * [687] 最长同值路径
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
func longestUnivaluePath(root *TreeNode) int {
	var res int
	// 递归函数功能：搜寻以node为起点的最长同值路径:要么是以node为起点的左子树，要么是以node为起点的右子树
	var longestPath func(node *TreeNode) int
	longestPath = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		length := 0
		left := longestPath(node.Left)
		right := longestPath(node.Right)
		// 这种情况对于寻找全局最长同值路径有帮助，对于搜索以node为路径起始点的最长同值路径没有帮助
		if node.Left != nil && node.Right != nil && node.Left.Val == node.Val && node.Right.Val == node.Val {
			res = max(res, left+right+2)
		}
		// 从左右子树中选择以node为路径起始点的最长同值路径
		if node.Left != nil && node.Left.Val == node.Val {
			length = left+1
		}
		if node.Right != nil && node.Right.Val == node.Val {
			length = max(length, right+1)
		}
		// 更新res	
		res = max(res,length)
		return length
	}
	longestPath(root)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end


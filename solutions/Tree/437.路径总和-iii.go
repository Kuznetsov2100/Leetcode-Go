/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
// 以当前节点作为根结点的子树的合法路径数量
// 以当前节点的左节点作为根节点的子树的合法路径数量
// 以当前节点的右节点作为根节点的子树的合法路径数量
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return countpath(root,sum) + pathSum(root.Left,sum) + pathSum(root.Right, sum)
}

func countpath(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum -= root.Val
	res := 0
	if 	sum == 0 {
		res = 1
	}
	return res+countpath(root.Left,sum)+countpath(root.Right,sum)
}
// @lc code=end


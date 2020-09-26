/*
 * @lc app=leetcode.cn id=1448 lang=golang
 *
 * [1448] 统计二叉树中好节点的数目
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
func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return  1 + helper(root.Left, root.Val) + helper(root.Right, root.Val)
}

func helper(root *TreeNode,minimum int) int {
	if root == nil {
		return 0
	}
	count := 0
	if root.Val >= minimum {
		count = 1
		minimum = root.Val
	}
	return count + helper(root.Left, minimum) + helper(root.Right, minimum)
}
// @lc code=end


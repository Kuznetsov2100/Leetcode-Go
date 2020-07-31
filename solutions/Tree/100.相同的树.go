/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil { // 节点都为空
		return true
	}
	if p == nil || q == nil { // 有一个节点为空
		return false
	}
	if p.Val != q.Val { // 节点的值不相等
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) // 节点值相等，接下来判断左子树和右子树是否都相等
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode)  {
	// 依次遍历6，5，4，3，2，1，其顺序为变形后的后序遍历
	// 遍历到一个节点时，将当前节点的右指针指向上一个节点，同时删除当前节点的左指针。
	// 用pre来保存上一个节点
	var pre *TreeNode
	var flat func(root *TreeNode)
	flat = func(root *TreeNode) {
		if root == nil {
			return
		}
		flat(root.Right)
		flat(root.Left)
		root.Right = pre
		root.Left = nil
		pre = root
	}
	flat(root)	
}
// @lc code=end


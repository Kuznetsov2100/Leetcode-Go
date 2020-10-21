/*
 * @lc app=leetcode.cn id=897 lang=golang
 *
 * [897] 递增顺序查找树
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
func increasingBST(root *TreeNode) *TreeNode {
	dummyHead := &TreeNode{}
	cur := dummyHead
	var inorder func(x *TreeNode)
	inorder = func(x *TreeNode) {
		if x == nil {
			return
		}
		inorder(x.Left)
		x.Left = nil
		cur.Right = x
		cur = x
		inorder(x.Right)

	}
	inorder(root)
	return dummyHead.Right
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
// 递归结束条件：
// 	都为空指针则返回 true
// 	只有一个为空则返回 false
// 递归过程：
// 判断两个指针当前节点值是否相等:
// 	判断 A 的右子树与 B 的左子树是否对称
// 	判断 A 的左子树与 B 的右子树是否对称
func isSymmetric(root *TreeNode) bool {
	return check(root,root)
}

func check(x,y *TreeNode) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	if x.Val != y.Val {
		return false
	}
	return check(x.Left,y.Right) && check(x.Right, y.Left)
}
// @lc code=end


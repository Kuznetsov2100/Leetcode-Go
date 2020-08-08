/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < min(p.Val,q.Val) {
		return lowestCommonAncestor(root.Right, p, q)
	} else if root.Val > max(p.Val,q.Val) {
		return lowestCommonAncestor(root.Left,p, q)
	} else {
		return root
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end


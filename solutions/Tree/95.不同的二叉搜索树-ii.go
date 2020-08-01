/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
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
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1,n)
}

func helper(lo,hi int) (res []*TreeNode) {
	if lo > hi {
		return []*TreeNode{nil} // nil是res中的一个空节点
	}
	// 枚举可行根节点
	for i := lo; i <= hi; i++ {
		leftPC := helper(lo, i-1) // [lo..i-1]为合法的左子树集合
		rightPC := helper(i+1,hi) // [i+1..hi]为合法的右子树集合
		// 从左子树集合中挑出一颗左子树，从右子树集合中挑出一颗右子树，拼接到以i为根节点的二叉搜索树中
		for _, leftTree := range leftPC {
			for _, rightTree := range rightPC {
				res = append(res, &TreeNode{Val: i, Left: leftTree, Right: rightTree})
			}
		}
	}
	return res
}
// @lc code=end


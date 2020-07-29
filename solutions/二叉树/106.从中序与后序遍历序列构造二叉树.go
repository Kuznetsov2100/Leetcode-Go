/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	// postorder最后一个元素即为根节点
	root := &TreeNode{Val:postorder[len(postorder)-1]}
	i := 0
	// 根据postorder的最后一个元素找到inorder的根节点在inorder数组中的索引
	for ; i < len(postorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			break			
		}
	}
	root.Left = buildTree(inorder[:i],postorder[:len(inorder[:i])]) // 递归构建左子树
	root.Right = buildTree(inorder[i+1:],postorder[len(inorder[:i]):len(postorder)-1]) // 递归构建右子树
	return root
}
// @lc code=end


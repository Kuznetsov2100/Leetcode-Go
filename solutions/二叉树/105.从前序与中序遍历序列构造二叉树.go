/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// preorder的一个元素即为根节点
	root := &TreeNode{Val:preorder[0]}
	i := 0
	// 根据preorder的第一个元素找到inorder的根节点在inorder数组中的索引
	for ; i < len(preorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1],inorder[:i]) // 递归构建左子树
	root.Right = buildTree(preorder[len(inorder[:i])+1:],inorder[i+1:]) // 递归构建右子树
	return root
}
// @lc code=end


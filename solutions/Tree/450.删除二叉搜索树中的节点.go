/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left,key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Right == nil {
			return root.Left
		}
		if root.Left == nil {
			return root.Right
		}
		t := root // 用变量t存储待删除的节点root
		root = min(t.Right) // 将root指向它的后继节点(二叉搜索树中为右子树中的最小节点)
		// 将root的右链接（原本指向一棵所有结点都大于min(t.Right).Val的二叉查找树指向
		//deleteMin(t.right)，也就是在删除后所有结点仍然都大于root.Val的子二叉查找树；
		root.Right = deleteNode(t.Right, root.Val)
		// 将root的左链接（本为空）设为t.Left
	    //（其下所有的键都小于被删除的结点和它的后继结点）。
		root.Left = t.Left
	}
	return root
}

// If the left link of the root is null, the smallest key in a BST is the key at the root;
// if the left link is not null, the smallest key in the BST is the smallest key in the subtree
// rooted at the node referenced by the left link.
func min(x *TreeNode) *TreeNode {
	if x.Left == nil {
		return x
	}
	return min(x.Left)
}
// @lc code=end


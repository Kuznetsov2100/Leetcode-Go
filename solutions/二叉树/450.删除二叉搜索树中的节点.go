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

// 1.Save a link to the node to be deleted in t
// 2.Set root to point to its successor min(t.right).
// 3.Set the right link of root (which is supposed to point to the BST containing all the keys larger than root.Val)
//   to deleteMin(t.right), the link to the BST containing all the keys that are larger than root.Val after the deletion.
// 4.Set the left link of root (which was null) to t.Left
//   (all the keys that are less than both the deleted key and its successor).
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
		t := root
		root = min(t.Right)
		root.Right = deleteMin(t.Right)
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

// For delete the minimum,
// we go left until finding a node that that has a null left link
// and then replace the link to that node by its right link.
func deleteMin(x *TreeNode) *TreeNode {
	if x.Left == nil {
		return x.Right
	}
	x.Left = deleteMin(x.Left)
	return x
}
// @lc code=end


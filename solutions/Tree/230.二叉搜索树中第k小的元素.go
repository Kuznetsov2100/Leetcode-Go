/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
func kthSmallest(root *TreeNode, k int) int {
	// Notice: k is 1-based, but the algorithm we used below is 0-based.
	// If the number of keys leftsize in the left subtree is larger than k-1,
	// we look (recursively) for the key of rank k in the left subtree;
	// if leftsize is equal to k-1, we return the key at the root;
	// and if leftsize is smaller than k-1,
	// we look (recursively) for the key of rank k - leftsizes - 1 in the right subtree.
	if leftSize := size(root.Left); leftSize > k-1 {
		return kthSmallest(root.Left, k)
	} else if leftSize < k-1 {
		return kthSmallest(root.Right, k-leftSize-1)
	} else {
		return root.Val
	}
}

func size(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return size(root.Left) + 1 + size(root.Right)
}
// @lc code=end


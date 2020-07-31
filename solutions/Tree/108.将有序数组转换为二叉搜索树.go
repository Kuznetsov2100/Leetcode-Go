/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
func sortedArrayToBST(nums []int) *TreeNode {
	return buildTree(nums, 0, len(nums)-1)
}

func buildTree(nums []int, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}
	mid := lo + (hi-lo)/2
	root := &TreeNode{Val:nums[mid]} // 以数组中点作为根节点
	root.Left = buildTree(nums,lo, mid-1) // 递归构建左子树
	root.Right = buildTree(nums,mid+1, hi) // 递归构建右子树
	return root
}
// @lc code=end


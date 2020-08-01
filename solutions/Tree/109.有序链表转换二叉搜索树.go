/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var array []int
	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}
	var buildTree func(lo,hi int) *TreeNode
	buildTree = func(lo, hi int) *TreeNode {
		if lo > hi {
			return nil
		}
		mid := lo + (hi-lo)/2
		return &TreeNode{
			Val: array[mid],
			Left: buildTree(lo, mid-1),
			Right: buildTree(mid+1, hi),
		}
	}
	return buildTree(0, len(array)-1)
}
// @lc code=end


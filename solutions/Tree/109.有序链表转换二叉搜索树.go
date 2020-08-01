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
	return buildTree(array,0, len(array)-1)
}

func buildTree(nums []int, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}
	mid := lo + (hi-lo)/2
	return &TreeNode{
		Val: nums[mid],
		Left: buildTree(nums,lo, mid-1),
		Right: buildTree(nums,mid+1, hi),
	}
}
// @lc code=end


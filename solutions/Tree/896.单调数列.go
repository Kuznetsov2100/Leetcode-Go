/*
 * @lc app=leetcode.cn id=896 lang=golang
 *
 * [896] 单调数列
 */

// @lc code=start
import "sort"
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isMonotonic(A []int) bool {
	B := make([]int, len(A))
	copy(B,A)
	for i, j := 0, len(A)-1; i < j; i, j = i+1, j-1 {
		A[i], A[j] = A[j], A[i]
	}
	if sort.IntsAreSorted(A) {
		tree := sortedArrayToBST(A)
		return isValidBST(tree)
	} else {
		tree := sortedArrayToBST(B)
		return isValidBST(tree)
	}
}


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

func isValidBST(root *TreeNode) bool {
	return helper(root, nil,nil)
}

func helper(root,min,max *TreeNode) bool {
	// root需要对整棵左子树和右子树进行约束
	if root == nil {
		return true
	}
	if min != nil && root.Val < min.Val {
		return false
	}
	if max != nil && root.Val > max.Val {
		return false
	}
	return helper(root.Left,min,root) && helper(root.Right,root,max)
}
// @lc code=end


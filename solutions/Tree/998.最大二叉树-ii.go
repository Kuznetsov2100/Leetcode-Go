/*
 * @lc app=leetcode.cn id=998 lang=golang
 *
 * [998] 最大二叉树 II
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
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	A := inorderTraversal(root)
	A = append(A, val)
	return constructMaximumBinaryTree(A)
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var inorder func(x *TreeNode)
	inorder = func(x *TreeNode)	{
		if x == nil {
			return
		}
		inorder(x.Left)
		res = append(res, x.Val)
		inorder(x.Right)
	}
	inorder(root)
	return res
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return buildTree(nums, 0, len(nums)-1)
}

func buildTree(nums []int, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}
	// 找到nums[lo..hi]中的最大值及最大值的索引
	val, index := func(nums []int, lo, hi int) (int,int) {
		max := -1 << 63
		index := 0
		for i := lo; i <= hi; i++ {
			if nums[i] > max  {
				max = nums[i]
				index = i
			}
		}
		return max,index
	}(nums,lo,hi)
	return &TreeNode{
		Val : val,
		Left: buildTree(nums,lo, index-1), // 递归构建左子树
		Right: buildTree(nums,index+1, hi), // 递归构建右子树
	}
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=1305 lang=golang
 *
 * [1305] 两棵二叉搜索树中的所有元素
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
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1 := inorderTraversal(root1)
	arr2 := inorderTraversal(root2)

	left, right, index := 0, 0, 0
	m, n := len(arr1), len(arr2)
	res := make([]int, m+n)
	for left < m && right < n {
		if arr1[left] < arr2[right] {
			res[index] = arr1[left]
			left++
		} else {
			res[index] = arr2[right]
			right++
		}
		index++
	}
	for left < m {
		res[index] = arr1[left]
		left++
		index++
	}
	for right < n {
		res[index] = arr2[right]
		right++
		index++
	}
	return res
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
// @lc code=end


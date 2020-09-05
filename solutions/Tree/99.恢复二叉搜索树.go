/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
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
func recoverTree(root *TreeNode)  {
	// 首先中序遍历，然后找到被错误交换的位置
	// 错误的位置有两种情况
	// case1： 错误的位置相邻，那么只有这相邻的两个位置不满足值有序性
	// case2： 错误的位置不相邻，那么有两对位置不满足值有序性
	arr := inorderTraversal(root)
	x, y := -1, -1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1].Val < arr[i].Val {
			y = i+1 // y可能需要更新
			if x == -1 {
				x = i // 找到x
			} else {
				break
			}
		}
	}
	arr[x].Val, arr[y].Val = arr[y].Val, arr[x].Val
}

func inorderTraversal(root *TreeNode) []*TreeNode {
	var res []*TreeNode
	var inorder func(x *TreeNode)
	inorder = func(x *TreeNode)	{
		if x == nil {
			return
		}
		inorder(x.Left)
		res = append(res, x)
		inorder(x.Right)
	}
	inorder(root)
	return res
}
// @lc code=end


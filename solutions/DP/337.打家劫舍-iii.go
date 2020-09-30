/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
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
 // 树形dp
 // 一个节点可以选择偷或不偷， 每个节点定义两个状态：
 // 当前节点选择偷，以该节点为根节点的子树偷得的最大金额 = 左子节点不偷得到的最大金额+右子节点不偷得到的最大金额
 // 当前节点不偷， 以该节点为根节点的子树偷得的最大金额 = 左子节点的最大金额+右子节点的最大金额
 // 采用后序遍历自底向上求出根节点的最大值 
func rob(root *TreeNode) int {
	return max(helper(root))
}

func helper(root *TreeNode) (undo int, do int) {
	if root == nil {
		return
	}
	leftUndo, leftDo := helper(root.Left)
	rightUndo, rightDo := helper(root.Right)
	return max(leftUndo, leftDo)+max(rightUndo, rightDo), root.Val+leftUndo+rightUndo
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 版本一： 备忘录  + 递归
// func rob(root *TreeNode) int {
// 	m := make(map[*TreeNode]int)
// 	var helper func(x *TreeNode) int
// 	helper = func(x *TreeNode) int {
// 		if x == nil {
// 			return 0
// 		}
// 		if val, ok := m[x]; ok {
// 			return val
// 		}
// 		robroot := x.Val
// 		if x.Left != nil {
// 			robroot += helper(x.Left.Left) + helper(x.Left.Right)
// 		}
// 		if x.Right != nil {
// 			robroot += helper(x.Right.Left) + helper(x.Right.Right)
// 		}
// 		notrob := helper(x.Left) + helper(x.Right)
// 		m[x] = max(robroot, notrob)
// 		return m[x]
// 	}
// 	return helper(root)
// }
// @lc code=end


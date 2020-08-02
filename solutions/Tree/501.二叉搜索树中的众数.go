/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
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
func findMode(root *TreeNode) []int {
//  使用中序遍历，遍历结果为非递减的序列
// 	如果pre不存在，说明目前是在第一个元素的位置，所以将计数置为1，并添加当前元素进入结果数组
//  如果pre存在，分两种情况计数:
//  如果pre.Val == root.Val：说明当前值和前一个相同，计数加1,count++
//  如果pre.Val == root.Val：说明当前值和前一个不同，计数置为1,count = 1
//  在得到计数count后，与当前最大计数maxnum比较：
//  如果count > maxcount: 更新maxcount, root.Val是结果数组中唯一的众数
//  如果count == maxcount: 将root.Val作为新的众数添加到结果数组中
	var res []int
	count, maxcount := 0,0
	var pre *TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre != nil && pre.Val == root.Val {
			count++
		} else {
			count = 1
		}
		if count > maxcount {
			maxcount = count
			res = []int{root.Val}
		} else if count == maxcount {
			res = append(res, root.Val)
		}
		pre = root
		dfs(root.Right)
	}
	dfs(root)
	return res
}
// @lc code=end


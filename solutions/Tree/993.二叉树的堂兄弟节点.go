/*
 * @lc app=leetcode.cn id=993 lang=golang
 *
 * [993] 二叉树的堂兄弟节点
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
func isCousins(root *TreeNode, x int, y int) bool {
	// bfs 层序遍历
	queue := []*TreeNode{root} // root节点加入队列
	for len(queue) != 0 {
		foundX, foundY := false, false
		for n := len(queue); n > 0; n-- {
			node := queue[0]
			queue = queue[1:]
			// 假如node的两个子节点含有x,y 直接返回false
			if node.Left != nil && node.Right != nil {
				if (node.Left.Val == x && node.Right.Val == y) ||
				 (node.Left.Val == y && node.Right.Val == x) {
					return false
				}
			}
			if node.Val == x { foundX = true } // 找到x
			if node.Val == y { foundY = true } // 找到y
			if node.Left != nil { queue = append(queue, node.Left) }
			if node.Right != nil { queue = append(queue, node.Right) }
		}
		if foundX && foundY { return true } // 在该层同时找到x,y
		if foundX || foundY { return false } // 在该层只找到了x,y中的一个，说明x,y不可能出现在同一层,直接返回false
	}
	return false
}
// @lc code=end


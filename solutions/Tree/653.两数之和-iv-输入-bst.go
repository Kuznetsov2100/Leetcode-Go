/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] 两数之和 IV - 输入 BST
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
// func findTarget(root *TreeNode, k int) bool {
// 	set := make(map[int]int)
// 	return find(root,k, set)
// }

// func find(root *TreeNode, k int, set map[int]int) bool {
// 	if root == nil {
// 		return false
// 	}
// 	if _, ok := set[k-root.Val]; ok {
// 		return true
// 	}
// 	set[root.Val] = 0
// 	return find(root.Left, k, set) || find(root.Right, k ,set)
// }
func findTarget(root *TreeNode, k int) bool {
	var queue []*TreeNode
	set := make(map[int]int)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if _, ok := set[k-node.Val]; ok {
			return true
		}
		set[node.Val] = 0
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return false
}

// @lc code=end


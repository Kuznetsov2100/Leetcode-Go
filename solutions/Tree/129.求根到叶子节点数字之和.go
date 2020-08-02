/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
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
import "strconv"
func sumNumbers(root *TreeNode) int {
	var sum int
	var dfs func(root *TreeNode, path string)
	dfs = func(root *TreeNode, path string) {
		if root != nil {
			path += strconv.Itoa(root.Val)
			if root.Left == nil && root.Right == nil {
				num, _ := strconv.Atoi(path)
				sum += num
			} else {
				dfs(root.Left, path)
				dfs(root.Right, path)
			}
		}
	}
	dfs(root, "")
	return sum
}
// @lc code=end


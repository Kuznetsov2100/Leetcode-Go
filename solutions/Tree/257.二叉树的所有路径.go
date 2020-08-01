/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	var dfs func(root *TreeNode, path string)
	dfs = func(root *TreeNode, path string) {
		if root != nil {
			path += strconv.Itoa(root.Val)
			if root.Left == nil && root.Right == nil {
				paths = append(paths, path) // 叶节点，将根节点到叶节点的路径加入到结果中
			} else {
				path += "->"
				dfs(root.Left, path)
				dfs(root.Right, path)
			}
		}
	}
	dfs(root, "")
	return paths
}
// @lc code=end

